import { useState } from 'react'
import axios from 'axios'
import './CreateUserForm.css'

export default function CreateUserForm({ roleId, title }) {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        phone_number: ''
    })
    const [status, setStatus] = useState(null)

    const handleSubmit = async (e) => {
        e.preventDefault()
        try {
            await axios.post(
                'http://localhost:8080/api/auth/sign-up-privileged',
                { ...formData, role_id: roleId },
                {
                    headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
                }
            )
            setStatus({ type: 'success', message: 'Пользователь успешно создан' })
            setFormData({ username: '', email: '', password: '', phone_number: '' })
        } catch (err) {
            setStatus({
                type: 'error',
                message: err.response?.data?.message || 'Ошибка при создании пользователя'
            })
        }
    }

    const handleChange = (e) => {
        setFormData(prev => ({
            ...prev,
            [e.target.name]: e.target.value
        }))
    }

    return (
        <div className="create-user-form">
            <h2>{title}</h2>
            {status && (
                <div className={`status-message ${status.type}`}>
                    {status.message}
                </div>
            )}
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="username">Имя пользователя</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="email">Email</label>
                    <input
                        type="email"
                        id="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Пароль</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="phone_number">Номер телефона</label>
                    <input
                        type="tel"
                        id="phone_number"
                        name="phone_number"
                        value={formData.phone_number}
                        onChange={handleChange}
                        required
                    />
                </div>
                <button type="submit">Создать</button>
            </form>
        </div>
    )
}
