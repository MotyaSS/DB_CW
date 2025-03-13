import { useState, useEffect } from 'react'
import axios from 'axios'
import './StoreDeleteForm.css'

export default function StoreDeleteForm({ onSubmit }) {
    const [stores, setStores] = useState([])
    const [selectedStore, setSelectedStore] = useState('')
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        fetchStores()
    }, [])

    const fetchStores = async () => {
        try {
            const response = await axios.get('http://localhost:8080/api/stores')
            setStores(response.data || [])
        } catch (err) {
            setError('Ошибка при загрузке магазинов')
        } finally {
            setLoading(false)
        }
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        if (selectedStore) {
            onSubmit(selectedStore)
        }
    }

    if (loading) return <div>Загрузка...</div>
    if (error) return <div className="error-message">{error}</div>

    return (
        <form className="store-delete-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="store">Выберите магазин</label>
                <select
                    id="store"
                    value={selectedStore}
                    onChange={(e) => setSelectedStore(e.target.value)}
                    required
                >
                    <option value="">Выберите магазин</option>
                    {stores.map(store => (
                        <option key={store.store_id} value={store.store_id}>
                            {store.store_name} ({store.store_address})
                        </option>
                    ))}
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