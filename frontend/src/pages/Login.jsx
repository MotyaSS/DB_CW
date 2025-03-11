import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import AuthForm from '../components/AuthForm/AuthForm'
import { useAuth } from '../context/AuthContext'
import './Auth.css'

export default function Login() {
    const [error, setError] = useState(null)
    const navigate = useNavigate()
    const { login } = useAuth()

    const handleLogin = async (formData) => {
        try {
            await login(formData.username, formData.password)
            navigate('/')
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при входе')
        }
    }

    return (
        <div className="auth-page">
            <h1>Вход</h1>
            {error && <div className="auth-error">{error}</div>}
            <AuthForm type="login" onSubmit={handleLogin} />
            <p className="auth-link">
                Нет аккаунта? <Link to="/register">Зарегистрируйтесь</Link>
            </p>
        </div>
    )
} 