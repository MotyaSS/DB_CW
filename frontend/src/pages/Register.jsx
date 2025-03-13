import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'
import './Register.css'

export default function Register() {
    const navigate = useNavigate()
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        phone_number: ''
    })
    const [error, setError] = useState(null)

    const formatPhoneNumber = (value) => {
        // Удаляем все нецифровые символы, кроме +
        const cleaned = value.replace(/[^\d+]/g, '')

        // Если строка пустая, возвращаем +7
        if (cleaned === '') return '+7'

        // Если первый символ не +, добавляем его
        if (!cleaned.startsWith('+')) {
            return '+7' + cleaned.substring(0, 10)
        }

        // Если начинается с +7, ограничиваем длину
        if (cleaned.startsWith('+7')) {
            return cleaned.substring(0, 12)
        }

        // В других случаях возвращаем +7
        return '+7'
    }

    const handleSubmit = async (e) => {
        e.preventDefault()

        // Проверяем формат номера телефона
        if (!/^\+7\d{10}$/.test(formData.phone_number)) {
            setError('Номер телефона должен быть в формате +7XXXXXXXXXX')
            return
        }

        try {
            await axios.post('http://localhost:8080/api/auth/sign-up', {
                username: formData.username,
                email: formData.email,
                password: formData.password,
                phone_number: formData.phone_number
            })
            navigate('/login')
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при регистрации')
        }
    }

    const handleChange = (e) => {
        if (e.target.name === 'phone_number') {
            setFormData(prev => ({
                ...prev,
                [e.target.name]: formatPhoneNumber(e.target.value)
            }))
            return
        }

        setFormData(prev => ({
            ...prev,
            [e.target.name]: e.target.value
        }))
    }

    return (
        <div className="register-container">
            <form className="register-form" onSubmit={handleSubmit}>
                <h2>Регистрация</h2>
                {error && <div className="error-message">{error}</div>}

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
                    <label htmlFor="phone_number">Номер телефона</label>
                    <input
                        type="tel"
                        id="phone_number"
                        name="phone_number"
                        value={formData.phone_number}
                        onChange={handleChange}
                        placeholder="+7"
                        minLength={12}
                        maxLength={12}
                        required
                    />
                    <small className="input-hint">Формат: +7XXXXXXXXXX</small>
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

                <button type="submit">Зарегистрироваться</button>

                <p className="login-link">
                    Уже есть аккаунт? <Link to="/login">Войти</Link>
                </p>
            </form>
        </div>
    )
} 