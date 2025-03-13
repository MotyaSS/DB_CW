import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import axios from 'axios'
import Modal from '../components/Modal/Modal'
import RentalForm from '../components/RentalForm/RentalForm'
import './InstrumentDetails.css'
import { formatPrice } from '../utils/formatters'
import Review from '../components/Review/Review'

export default function InstrumentDetails() {
    const { id } = useParams()
    const navigate = useNavigate()
    const { user } = useAuth()
    const [instrument, setInstrument] = useState(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [isRentalModalOpen, setRentalModalOpen] = useState(false)
    const defaultImage = 'https://placehold.co/600x400/1a1f2e/646cff?text=Нет+изображения'

    useEffect(() => {
        fetchInstrument()
    }, [id])

    const fetchInstrument = async () => {
        try {
            const response = await axios.get(`http://localhost:8080/api/instruments/${id}`)
            const instrumentData = response.data.Instrument || response.data
            setInstrument({
                ...instrumentData,
                discount: response.data.Discount
            })
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке инструмента')
        } finally {
            setLoading(false)
        }
    }

    const handleRent = async (rentalData) => {
        if (!user) {
            navigate('/login')
            return
        }

        try {
            await axios.post(
                `http://localhost:8080/api/instruments/${id}/rent`,
                rentalData,
                {
                    headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
                }
            )
            setRentalModalOpen(false)
            navigate('/profile')
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при аренде инструмента')
        }
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>
    if (!instrument) return <div className="error-message">Инструмент не найден</div>

    console.log('Instrument data:', instrument)

    return (
        <div className="instrument-details">
            <div className="instrument-details-content">
                <div className="instrument-image-large">
                    <img
                        src={instrument.image_url || defaultImage}
                        alt={instrument.instrument_name}
                        onError={(e) => {
                            e.target.src = defaultImage
                        }}
                    />
                </div>
                <div className="instrument-info-detailed">
                    <h1>{instrument.instrument_name}</h1>
                    <div className="price-section">
                        <p className="price">
                            {formatPrice(instrument.price_per_day || 0)} ₽/день
                        </p>
                        {instrument.discount && (
                            <div className="discount">
                                Скидка: {instrument.discount.discount_percentage}%
                            </div>
                        )}
                    </div>
                    <button className="rent-button" onClick={() => setRentalModalOpen(true)}>
                        {user ? 'Арендовать' : 'Войдите, чтобы арендовать'}
                    </button>
                </div>
            </div>
            <div className="instrument-sections">
                <section className="description-section">
                    <h2>Описание</h2>
                    <div className="description-content">
                        {instrument.description || 'Описание отсутствует'}
                    </div>
                </section>

                <section className="reviews-section">
                    <h2>Отзывы</h2>
                    <div className="reviews-list">
                        {/* Заглушка для отзывов */}
                        <Review review={{
                            rating: 5,
                            review_text: 'Отличный инструмент! Очень доволен арендой.',
                            created_at: new Date(),
                            user_name: 'Иван П.'
                        }} />
                        <Review review={{
                            rating: 4,
                            review_text: 'Хороший инструмент, но есть небольшие царапины.',
                            created_at: new Date(Date.now() - 86400000),
                            user_name: 'Мария С.'
                        }} />
                    </div>
                </section>
            </div>

            <Modal
                isOpen={isRentalModalOpen}
                onClose={() => setRentalModalOpen(false)}
                title="Аренда инструмента"
            >
                <RentalForm
                    instrumentId={id}
                    onSubmit={handleRent}
                    onCancel={() => setRentalModalOpen(false)}
                />
            </Modal>
        </div>
    )
} 