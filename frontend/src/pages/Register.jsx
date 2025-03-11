import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import AuthForm from '../components/AuthForm/AuthForm'
import axios from 'axios'
import './Auth.css'

export default function Register() {
    const [error, setError] = useState(null)
    const navigate = useNavigate()

    const handleRegister = async (formData) => {
        try {
            await axios.post('http://localhost:8080/api/auth/sign-up', {
                email: formData.email,
                password: formData.password,
                username: formData.username,
                name: formData.name,
                surname: formData.surname,
                phone: formData.phone
            })

            // После успешной регистрации перенаправляем на страницу входа
            navigate('/login')
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при регистрации')
        }
    }

    return (
        <div className="auth-page">
            <h1>Регистрация</h1>
            {error && <div className="auth-error">{error}</div>}
            <AuthForm type="register" onSubmit={handleRegister} />
            <p className="auth-link">
                Уже есть аккаунт? <Link to="/login">Войдите</Link>
            </p>
        </div>
    )
} 