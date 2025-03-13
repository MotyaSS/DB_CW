import { useState } from 'react'
import './StoreForm.css'

export default function StoreForm({ onSubmit }) {
    const [formData, setFormData] = useState({
        store_name: '',
        store_address: '',
        phone_number: ''
    })

    const handleSubmit = (e) => {
        e.preventDefault()
        onSubmit(formData)
        setFormData({ store_name: '', store_address: '', phone_number: '' })
    }

    const handlePhoneChange = (e) => {
        let value = e.target.value.replace(/\D/g, '')
        if (value.length > 0) {
            value = '+' + value
            if (value.length > 2) {
                value = value.slice(0, 2) + ' (' + value.slice(2)
            }
            if (value.length > 7) {
                value = value.slice(0, 7) + ') ' + value.slice(7)
            }
            if (value.length > 12) {
                value = value.slice(0, 12) + '-' + value.slice(12)
            }
            if (value.length > 15) {
                value = value.slice(0, 15) + '-' + value.slice(15)
            }
        }
        setFormData(prev => ({ ...prev, phone_number: value }))
    }

    return (
        <form className="store-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="store_name">Название магазина</label>
                <input
                    type="text"
                    id="store_name"
                    value={formData.store_name}
                    onChange={(e) => setFormData(prev => ({ ...prev, store_name: e.target.value }))}
                    required
                />
            </div>
            <div className="form-group">
                <label htmlFor="store_address">Адрес</label>
                <input
                    type="text"
                    id="store_address"
                    value={formData.store_address}
                    onChange={(e) => setFormData(prev => ({ ...prev, store_address: e.target.value }))}
                    required
                />
            </div>
            <div className="form-group">
                <label htmlFor="phone_number">Телефон</label>
                <input
                    type="tel"
                    id="phone_number"
                    value={formData.phone_number}
                    onChange={handlePhoneChange}
                    placeholder="+7 (999) 999-99-99"
                    required
                />
            </div>
            <button type="submit">Добавить магазин</button>
        </form>
    )
}