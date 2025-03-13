import { useState, useEffect } from 'react'
import axios from 'axios'
import { formatPrice } from '../../utils/formatters'
import './RentalHistory.css'

export default function RentalHistory() {
    const [rentals, setRentals] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        fetchRentals()
    }, [])

    const fetchRentals = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/rentals', {
                headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
            })
            setRentals(response.data || [])
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке истории аренд')
        } finally {
            setLoading(false)
        }
    }

    const formatDate = (dateString) => {
        return new Date(dateString).toLocaleDateString('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        })
    }

    const defaultImage = 'https://placehold.co/600x400/1a1f2e/646cff?text=Нет+изображения'

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <div className="rental-history">
            <h2>История аренд</h2>
            {rentals.length === 0 ? (
                <p className="no-rentals">У вас пока нет аренд</p>
            ) : (
                <div className="rentals-list">
                    {rentals.map(rental => (
                        <div key={rental.rental_id} className="rental-item">
                            <div className="instrument-image">
                                <img
                                    src={rental.instrument?.image_url || defaultImage}
                                    alt={rental.instrument?.instrument_name}
                                    onError={(e) => { e.target.src = defaultImage }}
                                />
                            </div>
                            <div className="instrument-info">
                                <h3>{rental.instrument?.instrument_name}</h3>
                                <p className="instrument-description">
                                    {rental.instrument?.description}
                                </p>
                                <span className="rental-dates">
                                    {formatDate(rental.rental_date)} - {formatDate(rental.return_date)}
                                </span>
                                <p className="rental-price">
                                    {formatPrice(rental.instrument?.price_per_day)} ₽/день
                                </p>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    )
} 