import { useState, useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'
import InstrumentCard from '../components/InstrumentCard/InstrumentCard'
import InstrumentFilters from '../components/InstrumentFilters/InstrumentFilters'
import axios from 'axios'
import './Instruments.css'

export default function Instruments() {
    const [instruments, setInstruments] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [searchParams] = useSearchParams()

    useEffect(() => {
        fetchInstruments()
    }, [searchParams])

    const fetchInstruments = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/instruments', {
                params: Object.fromEntries(searchParams)
            })
            setInstruments(response.data.items || [])
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке инструментов')
        } finally {
            setLoading(false)
        }
    }

    const handleFilterChange = (filters) => {
        // Фильтры обновляются через URL параметры
        // Это вызовет повторный запрос через useEffect
    }

    if (loading) return <div className="loading">Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <div className="instruments-page">
            <h1>Доступные инструменты</h1>
            <div className="instruments-layout">
                <aside className="instruments-filters">
                    <InstrumentFilters onFilterChange={handleFilterChange} />
                </aside>
                <div className="instruments-grid">
                    {instruments.map(instrument => (
                        <InstrumentCard
                            key={instrument.instrument_id}
                            instrument={instrument}
                        />
                    ))}
                </div>
            </div>
        </div>
    )
} 