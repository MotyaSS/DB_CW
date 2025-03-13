import { useState, useEffect } from 'react'
import axios from 'axios'
import './RentalForm.css'

export default function RentalForm({ instrumentId, onSubmit, onCancel }) {
    const [dates, setDates] = useState({
        start_date: '',
        end_date: ''
    })
    const [existingRentals, setExistingRentals] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        fetchExistingRentals()
    }, [instrumentId])

    const fetchExistingRentals = async () => {
        try {
            const response = await axios.get(
                `http://localhost:8080/api/instruments/${instrumentId}/rent`,
                {
                    headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
                }
            )
            setExistingRentals(response.data || [])
        } catch (err) {
            setError('Не удалось загрузить существующие аренды')
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

    const handleSubmit = (e) => {
        e.preventDefault()
        // Убедимся, что даты в правильном формате YYYY-MM-DD
        const formattedDates = {
            start_date: dates.start_date,
            end_date: dates.end_date
        }
        onSubmit(formattedDates)
    }

    const isDateAvailable = (date) => {
        // Проверяем, не попадает ли дата в промежуток существующих аренд
        return !existingRentals.some(rental => {
            const rentalStart = new Date(rental.start_date)
            const rentalEnd = new Date(rental.end_date)
            const checkDate = new Date(date)
            return checkDate >= rentalStart && checkDate <= rentalEnd
        })
    }

    if (loading) return <div>Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <form className="rental-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="start_date">Дата начала аренды</label>
                <input
                    type="date"
                    id="start_date"
                    value={dates.start_date}
                    onChange={(e) => {
                        const date = e.target.value
                        setDates(prev => ({ ...prev, start_date: date }))
                    }}
                    min={new Date().toISOString().split('T')[0]}
                    required
                />
            </div>

            <div className="form-group">
                <label htmlFor="end_date">Дата окончания аренды</label>
                <input
                    type="date"
                    id="end_date"
                    value={dates.end_date}
                    onChange={(e) => setDates(prev => ({ ...prev, end_date: e.target.value }))}
                    min={dates.start_date || new Date().toISOString().split('T')[0]}
                    required
                />
            </div>

            {existingRentals.length > 0 && (
                <div className="existing-rentals">
                    <h4>Занятые даты:</h4>
                    <ul>
                        {existingRentals.map((rental, index) => (
                            <li key={index}>
                                {formatDate(rental.rental_date)} - {formatDate(rental.return_date)}
                            </li>
                        ))}
                    </ul>
                </div>
            )}

            <div className="form-actions">
                <button type="submit">Арендовать</button>
                <button type="button" onClick={onCancel} className="cancel-button">
                    Отмена
                </button>
            </div>
        </form>
    )
} 