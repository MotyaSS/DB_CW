import { useState, useEffect } from 'react'
import axios from 'axios'
import './InstrumentDeleteForm.css'

export default function InstrumentDeleteForm({ onSubmit }) {
    const [instruments, setInstruments] = useState([])
    const [selectedInstrument, setSelectedInstrument] = useState('')
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        fetchInstruments()
    }, [])

    const fetchInstruments = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/instruments')
            setInstruments(response.data.items || [])
        } catch (err) {
            setError('Ошибка при загрузке инструментов')
        } finally {
            setLoading(false)
        }
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        if (selectedInstrument) {
            onSubmit(selectedInstrument)
        }
    }

    if (loading) return <div>Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <form className="instrument-delete-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="instrument">Выберите инструмент</label>
                <select
                    id="instrument"
                    value={selectedInstrument}
                    onChange={(e) => setSelectedInstrument(e.target.value)}
                    required
                >
                    <option value="">Выберите инструмент</option>
                    {instruments.map(item => {
                        const instrument = item.Instrument || item
                        return (
                            <option key={instrument.instrument_id} value={instrument.instrument_id}>
                                {instrument.instrument_name}
                            </option>
                        )
                    })}
                </select>
            </div>
            <div className="form-actions">
                <button type="submit" className="delete-button">
                    Удалить
                </button>
            </div>
        </form>
    )
}