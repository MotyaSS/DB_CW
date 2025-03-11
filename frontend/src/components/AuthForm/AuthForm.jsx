import { useState } from 'react'
import './AuthForm.css'

export default function AuthForm({ type, onSubmit }) {
    const [formData, setFormData] = useState({
        username: '',
        password: '',
        name: '',
        surname: '',
        email: '',
        phone: ''
    })

    const handleSubmit = (e) => {
        e.preventDefault()
        onSubmit(formData)
    }

    const handleChange = (e) => {
        const { name, value } = e.target

        if (name === 'phone') {
            // Форматируем телефон
            let phoneNumber = value.replace(/\D/g, '')

            if (!phoneNumber.startsWith('7') && phoneNumber.length > 0) {
                phoneNumber = '7' + phoneNumber
            }

            if (phoneNumber.length > 11) {
                phoneNumber = phoneNumber.slice(0, 11)
            }

            // Форматируем для отображения
            if (phoneNumber.length > 0) {
                phoneNumber = '+' + phoneNumber
                if (phoneNumber.length > 2) {
                    phoneNumber = phoneNumber.slice(0, 2) + ' (' + phoneNumber.slice(2)
                }
                if (phoneNumber.length > 7) {
                    phoneNumber = phoneNumber.slice(0, 7) + ') ' + phoneNumber.slice(7)
                }
                if (phoneNumber.length > 12) {
                    phoneNumber = phoneNumber.slice(0, 12) + '-' + phoneNumber.slice(12)
                }
                if (phoneNumber.length > 15) {
                    phoneNumber = phoneNumber.slice(0, 15) + '-' + phoneNumber.slice(15)
                }
            }

            setFormData(prev => ({
                ...prev,
                [name]: phoneNumber
            }))
        } else {
            setFormData(prev => ({
                ...prev,
                [name]: value
            }))
        }
    }

    const handlePhoneFocus = (e) => {
        if (!formData.phone) {
            setFormData(prev => ({
                ...prev,
                phone: '+7'
            }))
        }
    }

    return (
        <form className="auth-form" onSubmit={handleSubmit}>
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
            {type === 'register' && (
                <>
                    <div className="form-group">
                        <label htmlFor="email">Email</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            placeholder="mail@example.com"
                            required
                        />
                        <span className="input-hint">Формат: mail@example.com</span>
                    </div>
                    <div className="form-group">
                        <label htmlFor="name">Имя</label>
                        <input
                            type="text"
                            id="name"
                            name="name"
                            value={formData.name}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="surname">Фамилия</label>
                        <input
                            type="text"
                            id="surname"
                            name="surname"
                            value={formData.surname}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="phone">Телефон</label>
                        <input
                            type="tel"
                            id="phone"
                            name="phone"
                            value={formData.phone}
                            onChange={handleChange}
                            onFocus={handlePhoneFocus}
                            placeholder="+7 (999) 999-99-99"
                            required
                        />
                        <span className="input-hint">Формат: +7 (999) 999-99-99</span>
                    </div>
                </>
            )}
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
            <button type="submit">
                {type === 'login' ? 'Войти' : 'Зарегистрироваться'}
            </button>
        </form>
    )
} 