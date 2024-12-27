import { useState, useEffect } from 'react'
import { apiClient } from '../api/client'
import InstrumentCard from '../components/InstrumentCard/InstrumentCard'
import './Instruments.css'

export default function Instruments() {
    const [instruments, setInstruments] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [filters, setFilters] = useState({
        search: '',
        category: '',
        status: '',
        priceRange: [0, 10000]
    })

    useEffect(() => {
        fetchInstruments()
    }, [])

    const fetchInstruments = async () => {
        try {
            setLoading(true)
            const response = await apiClient.get('/api/instruments')
            setInstruments(response.data)
        } catch (err) {
            setError('Не удалось загрузить инструменты')
            console.error('Error fetching instruments:', err)
        } finally {
            setLoading(false)
        }
    }

    const handleFilterChange = (e) => {
        const { name, value } = e.target
        setFilters(prev => ({
            ...prev,
            [name]: value
        }))
    }

    const filteredInstruments = instruments.filter(instrument => {
        return instrument.name.toLowerCase().includes(filters.search.toLowerCase()) &&
            (filters.category === '' || instrument.category === filters.category) &&
            (filters.status === '' || instrument.status === filters.status) &&
            instrument.price >= filters.priceRange[0] &&
            instrument.price <= filters.priceRange[1]
    })

    if (loading) return <div className="instruments-loading">Загрузка...</div>
    if (error) return <div className="instruments-error">{error}</div>

    return (
        <div className="instruments-page">
            <div className="instruments-filters">
                <input
                    type="text"
                    name="search"
                    placeholder="Поиск инструментов..."
                    value={filters.search}
                    onChange={handleFilterChange}
                    className="filter-input"
                />
                <select
                    name="category"
                    value={filters.category}
                    onChange={handleFilterChange}
                    className="filter-select"
                >
                    <option value="">Все категории</option>
                    <option value="string">Струнные</option>
                    <option value="wind">Духовые</option>
                    <option value="percussion">Ударные</option>
                </select>
                <select
                    name="status"
                    value={filters.status}
                    onChange={handleFilterChange}
                    className="filter-select"
                >
                    <option value="">Все статусы</option>
                    <option value="available">Доступные</option>
                    <option value="rented">Арендованные</option>
                </select>
            </div>

            <div className="instruments-grid">
                {filteredInstruments.length > 0 ? (
                    filteredInstruments.map(instrument => (
                        <InstrumentCard 
                            key={instrument.id} 
                            instrument={instrument}
                        />
                    ))
                ) : (
                    <div className="instruments-empty">
                        Инструменты не найдены
                    </div>
                )}
            </div>
        </div>
    )
} 