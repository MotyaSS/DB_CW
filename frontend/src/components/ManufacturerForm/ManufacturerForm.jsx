import { useState } from 'react'
import './ManufacturerForm.css'

export default function ManufacturerForm({ onSubmit }) {
    const [manufacturerName, setManufacturerName] = useState('')

    const handleSubmit = (e) => {
        e.preventDefault()
        onSubmit({ manufacturer_name: manufacturerName })
        setManufacturerName('')
    }

    return (
        <form className="manufacturer-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="manufacturer_name">Название производителя</label>
                <input
                    type="text"
                    id="manufacturer_name"
                    value={manufacturerName}
                    onChange={(e) => setManufacturerName(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Добавить производителя</button>
        </form>
    )
} 