import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import axios from 'axios'
import './InstrumentDetails.css'
import { formatPrice } from '../utils/formatters'

export default function InstrumentDetails() {
    const { id } = useParams()
    const navigate = useNavigate()
    const { user } = useAuth()
    const [instrument, setInstrument] = useState(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const defaultImage = 'https://via.placeholder.com/600x400?text=Нет+изображения'

    useEffect(() => {
        fetchInstrument()
    }, [id])

    const fetchInstrument = async () => {
        try {
            const response = await axios.get(`http://localhost:8080/api/instruments/${id}`)
            setInstrument(response.data)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке инструмента')
        } finally {
            setLoading(false)
        }
    }

    const handleRent = async () => {
        if (!user) {
            navigate('/login')
            return
        }

        try {
            await axios.post(
                `http://localhost:8080/api/instruments/${id}/rent`,
                {},
                {
                    headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
                }
            )
            navigate('/profile') // После успешной аренды перенаправляем в профиль
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при аренде инструмента')
        }
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>
    if (!instrument) return <div className="error-message">Инструмент не найден</div>

    return (
        <div className="instrument-details">
            <div className="instrument-details-content">
                <div className="instrument-image-large">
                    <img
                        src={instrument.image_url || defaultImage}
                        alt={instrument.instrument_name}
                        onError={(e) => { e.target.src = defaultImage }}
                    />
                </div>
                <div className="instrument-info-detailed">
                    <h1>{instrument.instrument_name}</h1>
                    <p className="description">{instrument.description}</p>
                    <div className="price-section">
                        <p className="price">{formatPrice(instrument.price_per_day)} ₽/день</p>
                        {instrument.discount && (
                            <div className="discount">
                                Скидка: {instrument.discount.discount_percentage}%
                            </div>
                        )}
                    </div>
                    <button className="rent-button" onClick={handleRent}>
                        {user ? 'Арендовать' : 'Войдите, чтобы арендовать'}
                    </button>
                </div>
            </div>
        </div>
    )
} 