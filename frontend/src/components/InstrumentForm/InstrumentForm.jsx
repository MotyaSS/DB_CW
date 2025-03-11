import { useState, useEffect } from 'react'
import axios from 'axios'
import './InstrumentForm.css'

export default function InstrumentForm({ onSubmit }) {
    const [categories, setCategories] = useState([])
    const [manufacturers, setManufacturers] = useState([])
    const [stores, setStores] = useState([])
    const [loading, setLoading] = useState(true)
    const [status, setStatus] = useState(null)

    const [formData, setFormData] = useState({
        instrument_name: '',
        category_id: '',
        manufacturer_id: '',
        store_id: '',
        description: '',
        price_per_day: '',
        image_url: ''
    })

    useEffect(() => {
        const fetchData = async () => {
            try {
                const [categoriesRes, manufacturersRes, storesRes] = await Promise.all([
                    axios.get('http://localhost:8080/api/instruments/categories'),
                    axios.get('http://localhost:8080/api/instruments/manufacturers'),
                    axios.get('http://localhost:8080/api/stores')
                ])
                setCategories(categoriesRes.data || [])
                setManufacturers(manufacturersRes.data || [])
                setStores(storesRes.data || [])
            } catch (err) {
                console.error('Error fetching data:', err)
            } finally {
                setLoading(false)
            }
        }
        fetchData()
    }, [])

    const handleChange = (e) => {
        const { name, value } = e.target
        setFormData(prev => ({
            ...prev,
            [name]: value
        }))
    }

    const handleImageChange = async (e) => {
        const file = e.target.files[0]
        if (!file) return

        const reader = new FileReader()
        reader.onloadend = () => {
            setImagePreview(reader.result)
        }
        reader.readAsDataURL(file)

        try {
            const formData = new FormData()
            formData.append('image', file)

            const response = await axios.post('https://api.imgur.com/3/image', formData, {
                headers: {
                    'Authorization': 'Client-ID YOUR_IMGUR_CLIENT_ID'
                }
            })

            setFormData(prev => ({
                ...prev,
                image_url: response.data.data.link
            }))
        } catch (err) {
            console.error('Error uploading image:', err)
        }
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const formattedData = {
            ...formData,
            category_id: parseInt(formData.category_id),
            manufacturer_id: parseInt(formData.manufacturer_id),
            store_id: parseInt(formData.store_id),
            price_per_day: parseFloat(formData.price_per_day)
        }

        try {
            await onSubmit(formattedData)
            setStatus({ type: 'success', message: 'Инструмент успешно добавлен' })
            setFormData({
                instrument_name: '',
                category_id: '',
                manufacturer_id: '',
                store_id: '',
                description: '',
                price_per_day: '',
                image_url: ''
            })
        } catch (err) {
            setStatus({
                type: 'error',
                message: err.response?.data?.message || 'Ошибка при добавлении инструмента'
            })
        }

        setTimeout(() => {
            setStatus(null)
        }, 3000)
    }

    if (loading) return <div>Загрузка...</div>

    return (
        <div className="instrument-form-container">
            {status && (
                <div className={`status-message ${status.type}`}>
                    {status.message}
                </div>
            )}
            <form className="instrument-form" onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="instrument_name">Название инструмента</label>
                    <input
                        type="text"
                        id="instrument_name"
                        name="instrument_name"
                        value={formData.instrument_name}
                        onChange={handleChange}
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="category_id">Категория</label>
                    <select
                        id="category_id"
                        name="category_id"
                        value={formData.category_id}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите категорию</option>
                        {categories.map(category => (
                            <option key={category.category_id} value={category.category_id}>
                                {category.category_name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="form-group">
                    <label htmlFor="manufacturer_id">Производитель</label>
                    <select
                        id="manufacturer_id"
                        name="manufacturer_id"
                        value={formData.manufacturer_id}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите производителя</option>
                        {manufacturers.map(manufacturer => (
                            <option key={manufacturer.manufacturer_id} value={manufacturer.manufacturer_id}>
                                {manufacturer.manufacturer_name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="form-group">
                    <label htmlFor="store_id">Магазин</label>
                    <select
                        id="store_id"
                        name="store_id"
                        value={formData.store_id}
                        onChange={handleChange}
                        required
                    >
                        <option value="">Выберите магазин</option>
                        {stores.map(store => (
                            <option key={store.store_id} value={store.store_id}>
                                {store.store_name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="form-group">
                    <label htmlFor="price_per_day">Цена за день</label>
                    <input
                        type="number"
                        id="price_per_day"
                        name="price_per_day"
                        value={formData.price_per_day}
                        onChange={handleChange}
                        min="0"
                        step="0.01"
                        required
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="description">Описание (необязательно)</label>
                    <textarea
                        id="description"
                        name="description"
                        value={formData.description}
                        onChange={handleChange}
                    />
                </div>

                <div className="form-group">
                    <label htmlFor="image_url">URL изображения</label>
                    <input
                        type="url"
                        id="image_url"
                        name="image_url"
                        value={formData.image_url}
                        onChange={handleChange}
                        placeholder="https://example.com/image.jpg"
                    />
                    {formData.image_url && (
                        <div className="image-preview">
                            <img
                                src={formData.image_url}
                                alt="Preview"
                                onError={(e) => e.target.style.display = 'none'}
                            />
                        </div>
                    )}
                </div>

                <button type="submit">Добавить инструмент</button>
            </form>
        </div>
    )
} 