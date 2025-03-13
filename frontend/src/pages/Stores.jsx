import { useState, useEffect } from 'react'
import StoreCard from '../components/StoreCard/StoreCard'
import axios from 'axios'
import './Stores.css'

export default function Stores() {
    const [stores, setStores] = useState([])
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
            setError(err.message)
        } finally {
            setLoading(false)
        }
    }

    if (loading) return <div>Загрузка...</div>
    if (error) return <div>Ошибка: {error}</div>

    return (
        <div className="stores-page">
            <h1>Наши магазины</h1>
            {stores.length === 0 ? (
                <div className="no-stores">
                    Магазины пока не добавлены
                </div>
            ) : (
                <div className="stores-grid">
                    {stores.map((store, index) => (
                        <StoreCard
                            key={store.id || `store-${index}`}
                            store={store}
                        />
                    ))}
                </div>
            )}
        </div>
    )
} 