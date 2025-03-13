import { useState } from 'react'
import './RepairForm.css'

export default function RepairForm({ instruments, onSubmit }) {
    const [formData, setFormData] = useState({
        instrument_id: '',
        repair_start_date: '',
        repair_end_date: '',
        repair_cost: '',
        description: ''
    })

    const handleSubmit = (e) => {
        e.preventDefault()
        onSubmit({
            ...formData,
            repair_cost: Number(formData.repair_cost)
        })
        setFormData({
            instrument_id: '',
            repair_start_date: '',
            repair_end_date: '',
            repair_cost: '',
            description: ''
        })
    }

    return (
        <form className="repair-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="instrument_id">Инструмент</label>
                <select
                    id="instrument_id"
                    value={formData.instrument_id}
                    onChange={(e) => setFormData(prev => ({ ...prev, instrument_id: e.target.value }))}
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

            <div className="form-group">
                <label htmlFor="repair_start_date">Дата начала ремонта</label>
                <input
                    type="date"
                    id="repair_start_date"
                    value={formData.repair_start_date}
                    onChange={(e) => setFormData(prev => ({ ...prev, repair_start_date: e.target.value }))}
                    required
                />
            </div>

            <div className="form-group">
                <label htmlFor="repair_end_date">Дата окончания ремонта</label>
                <input
                    type="date"
                    id="repair_end_date"
                    value={formData.repair_end_date}
                    onChange={(e) => setFormData(prev => ({ ...prev, repair_end_date: e.target.value }))}
                    required
                />
            </div>

            <div className="form-group">
                <label htmlFor="description">Описание ремонта</label>
                <textarea
                    id="description"
                    value={formData.description}
                    onChange={(e) => setFormData(prev => ({ ...prev, description: e.target.value }))}
                    required
                />
            </div>

            <div className="form-group">
                <label htmlFor="repair_cost">Стоимость ремонта (₽)</label>
                <input
                    type="number"
                    id="repair_cost"
                    value={formData.repair_cost}
                    onChange={(e) => setFormData(prev => ({ ...prev, repair_cost: e.target.value }))}
                    min="0"
                    step="0.01"
                    required
                />
            </div>

            <button type="submit">Добавить запись</button>
        </form>
    )
}
