import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import axios from 'axios'
import UserList from '../components/UserList/UserList'
import CreateUserForm from '../components/CreateUserForm/CreateUserForm'
import Modal from '../components/Modal/Modal'
import './Admin.css'

export default function Admin() {
    const { user } = useAuth()
    const [users, setUsers] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [isCreateUserModalOpen, setCreateUserModalOpen] = useState(false)
    const navigate = useNavigate()
    const token = localStorage.getItem('token')

    useEffect(() => {
        if (!token) {
            navigate('/login')
            return
        }
        fetchUsers()
    }, [token, navigate])

    const fetchUsers = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/users', {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            const userData = response.data.items || []
            setUsers(Array.isArray(userData) ? userData : [])
        } catch (err) {
            if (err.response?.status === 401) {
                navigate('/login')
            }
            setError(err.response?.data?.message || 'Ошибка при загрузке пользователей')
        } finally {
            setLoading(false)
        }
    }

    const handleDeleteUser = async (userId) => {
        if (!window.confirm('Вы уверены, что хотите удалить этого пользователя?')) {
            return
        }

        try {
            await axios.delete(`http://localhost:8080/api/users/${userId}`, {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })
            setUsers(users.filter(user => user.user_id !== userId))
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при удалении пользователя')
        }
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <div className="admin-page">
            <div className="admin-header">
                <h1>Панель администратора</h1>
                <button
                    className="create-user-btn"
                    onClick={() => setCreateUserModalOpen(true)}
                >
                    Создать аккаунт менеджера
                </button>
            </div>
            <div className="admin-content">
                <section className="users-section">
                    <h2>Пользователи</h2>
                    {users.length > 0 ? (
                        <UserList
                            users={users}
                            onDelete={handleDeleteUser}
                        />
                    ) : (
                        <p className="no-users">Пользователи не найдены</p>
                    )}
                </section>
            </div>
            <Modal
                isOpen={isCreateUserModalOpen}
                onClose={() => setCreateUserModalOpen(false)}
                title="Создать аккаунт менеджера"
            >
                <CreateUserForm
                    roleId={2}
                    title="Создать аккаунт менеджера"
                />
            </Modal>
        </div>
    )
} 