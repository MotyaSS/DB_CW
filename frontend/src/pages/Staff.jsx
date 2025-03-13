import { useState, useEffect } from 'react'
import axios from 'axios'
import Modal from '../components/Modal/Modal'
import RepairForm from '../components/RepairForm/RepairForm'
import { formatPrice } from '../utils/formatters'
import './Staff.css'

export default function Staff() {
    const [instruments, setInstruments] = useState([])
    const [repairs, setRepairs] = useState({})
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [success, setSuccess] = useState(null)
    const [isRepairModalOpen, setRepairModalOpen] = useState(false)
    const token = localStorage.getItem('token')

    useEffect(() => {
        fetchInstruments()
    }, [])

    const fetchInstruments = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/instruments')
            const instrumentsData = response.data.items || []
            setInstruments(instrumentsData)

            // Получаем ремонты для каждого инструмента
            const repairsData = {}
            await Promise.all(
                instrumentsData.map(async (item) => {
                    const instrument = item.Instrument
                    try {
                        const repairsResponse = await axios.get(
                            `http://localhost:8080/api/instruments/${instrument.instrument_id}/repairments`,
                            {
                                headers: { Authorization: `Bearer ${token}` }
                            }
                        )
                        repairsData[instrument.instrument_id] = repairsResponse.data || []
                    } catch (err) {
                        console.error(`Ошибка при загрузке ремонтов для инструмента ${instrument.instrument_id}:`, err)
                    }
                })
            )
            setRepairs(repairsData)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке данных')
        } finally {
            setLoading(false)
        }
    }

    const handleAddRepair = async (repairData) => {
        try {
            await axios.post(
                `http://localhost:8080/api/instruments/${repairData.instrument_id}/repairments`,
                {
                    repair_start_date: repairData.repair_start_date,
                    repair_end_date: repairData.repair_end_date,
                    repair_cost: Number(repairData.repair_cost),
                    description: repairData.description
                },
                {
                    headers: { Authorization: `Bearer ${token}` }
                }
            )
            setSuccess('Запись о ремонте успешно добавлена')
            setError(null)
            setRepairModalOpen(false)
            // Обновляем список ремонтов
            fetchInstruments()
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при добавлении записи о ремонте')
            setSuccess(null)
        }
    }

    const formatDate = (dateString) => {
        return new Date(dateString).toLocaleDateString('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        })
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <div className="staff-page">
            <h1>Панель сотрудника</h1>
            {success && <div className="success-message">{success}</div>}
            {error && <div className="error-message">{error}</div>}

            <div className="section-header">
                <h2>Ремонты инструментов</h2>
                <button
                    className="add-repair-button"
                    onClick={() => setRepairModalOpen(true)}
                >
                    Добавить запись о ремонте
                </button>
            </div>

            <div className="instruments-grid">
                {instruments.map(item => {
                    const instrument = item.Instrument
                    const instrumentRepairs = repairs[instrument.instrument_id] || []

                    return (
                        <div key={instrument.instrument_id} className="instrument-card">
                            <div className="instrument-info">
                                <h3>{instrument.instrument_name}</h3>
                                <p className="price">
                                    {formatPrice(instrument.price_per_day)} ₽/день
                                </p>
                            </div>

                            <div className="repairs-container">
                                <h4>История ремонтов:</h4>
                                {instrumentRepairs.length > 0 ? (
                                    <div className="repairs-list">
                                        {instrumentRepairs.map(repair => (
                                            <div key={repair.repair_id} className="repair-item">
                                                <div className="repair-dates">
                                                    {formatDate(repair.repair_start_date)} - {formatDate(repair.repair_end_date)}
                                                </div>
                                                <p className="repair-description">{repair.description}</p>
                                                <p className="repair-cost">
                                                    Стоимость: {formatPrice(repair.repair_cost)} ₽
                                                </p>
                                            </div>
                                        ))}
                                    </div>
                                ) : (
                                    <p className="no-repairs">Нет записей о ремонте</p>
                                )}
                            </div>
                        </div>
                    )
                })}
            </div>

            <Modal
                isOpen={isRepairModalOpen}
                onClose={() => setRepairModalOpen(false)}
                title="Добавить запись о ремонте"
            >
                <RepairForm
                    instruments={instruments}
                    onSubmit={handleAddRepair}
                />
            </Modal>
        </div>
    )
} 