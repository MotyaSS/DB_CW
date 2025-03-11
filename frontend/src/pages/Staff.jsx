import { useState, useEffect } from 'react'
import axios from 'axios'
import { formatPrice } from '../utils/formatters'
import './Staff.css'

export default function Staff() {
    const [instruments, setInstruments] = useState([])
    const [repairs, setRepairs] = useState({}) // Объект, где ключ - ID инструмента
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const token = localStorage.getItem('token')

    useEffect(() => {
        fetchInstruments()
    }, [])

    const fetchInstruments = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/instruments')
            console.log('Raw instruments data:', response.data)

            const instrumentsData = response.data.items.map(item => item.Instrument) || []
            setInstruments(instrumentsData)

            // Получаем ремонты для каждого инструмента
            const repairsData = {}
            await Promise.all(
                instrumentsData.map(async (instrument) => {
                    const repairsResponse = await axios.get(
                        `http://localhost:8080/api/instruments/${instrument.instrument_id}/repairments`,
                        {
                            headers: { Authorization: `Bearer ${token}` }
                        }
                    )
                    repairsData[instrument.instrument_id] = repairsResponse.data || []
                })
            )
            setRepairs(repairsData)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке данных')
        } finally {
            setLoading(false)
        }
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <div className="staff-page">
            <h1>Панель сотрудника</h1>
            <div className="staff-content">
                <section className="repairs-section">
                    <h2>Ремонты инструментов</h2>
                    <div className="instruments-list">
                        {instruments.map(instrument => (
                            <div key={instrument.instrument_id} className="instrument-repairs">
                                <div className="instrument-summary">
                                    <div className="instrument-info">
                                        <h3>{instrument.instrument_name}</h3>
                                        <div className="instrument-details">
                                            <p className="price">
                                                Цена аренды: {formatPrice(instrument.price_per_day)} ₽/день
                                            </p>
                                            {instrument.description && (
                                                <p className="description">{instrument.description}</p>
                                            )}
                                        </div>
                                    </div>
                                    {instrument.image_url && (
                                        <div className="instrument-image">
                                            <img
                                                src={instrument.image_url}
                                                alt={instrument.instrument_name}
                                                onError={(e) => e.target.style.display = 'none'}
                                            />
                                        </div>
                                    )}
                                </div>

                                <div className="repairs-container">
                                    <h4>История ремонтов</h4>
                                    {repairs[instrument.instrument_id]?.length > 0 ? (
                                        <div className="repairs-list">
                                            {repairs[instrument.instrument_id].map(repair => (
                                                <div key={repair.repair_id} className="repair-item">
                                                    <p className="repair-dates">
                                                        {new Date(repair.repair_start_date).toLocaleDateString()} -
                                                        {new Date(repair.repair_end_date).toLocaleDateString()}
                                                    </p>
                                                    <p className="repair-cost">
                                                        Стоимость: {formatPrice(repair.repair_cost)} ₽
                                                    </p>
                                                    <p className="repair-description">
                                                        {repair.description}
                                                    </p>
                                                </div>
                                            ))}
                                        </div>
                                    ) : (
                                        <p className="no-repairs">Нет записей о ремонте</p>
                                    )}
                                </div>
                            </div>
                        ))}
                    </div>
                </section>
            </div>
        </div>
    )
} 