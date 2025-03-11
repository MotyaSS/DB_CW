import { useState, useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'
import axios from 'axios'
import './InstrumentFilters.css'

export default function InstrumentFilters({ onFilterChange }) {
    const [searchParams, setSearchParams] = useSearchParams()
    const [categories, setCategories] = useState([])
    const [manufacturers, setManufacturers] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    const [filters, setFilters] = useState({
        category: searchParams.getAll('category') || [],
        manufacturer: searchParams.getAll('manufacturer') || [],
        price_floor: searchParams.get('price_floor') || '',
        price_ceil: searchParams.get('price_ceil') || '',
        page: searchParams.get('page') || 1
    })

    useEffect(() => {
        const fetchFilterData = async () => {
            try {
                const [categoriesRes, manufacturersRes] = await Promise.all([
                    axios.get('http://localhost:8080/api/instruments/categories'),
                    axios.get('http://localhost:8080/api/instruments/manufacturers')
                ])
                setCategories(categoriesRes.data || [])
                setManufacturers(manufacturersRes.data || [])
            } catch (err) {
                setError(err.message)
                setCategories([])
                setManufacturers([])
            } finally {
                setLoading(false)
            }
        }

        fetchFilterData()
    }, [])

    const handleFilterChange = (e) => {
        const { name, value } = e.target
        let newFilters = { ...filters }

        if (name === 'category' || name === 'manufacturer') {
            const values = [...newFilters[name]]
            const index = values.indexOf(value)
            if (index === -1) {
                values.push(value)
            } else {
                values.splice(index, 1)
            }
            newFilters[name] = values
        } else {
            newFilters[name] = value
        }

        setFilters(newFilters)
        updateSearchParams(newFilters)
        onFilterChange(newFilters)
    }

    const updateSearchParams = (filters) => {
        const params = new URLSearchParams()

        filters.category.forEach(cat => params.append('category', cat))
        filters.manufacturer.forEach(man => params.append('manufacturer', man))

        if (filters.price_floor) params.set('price_floor', filters.price_floor)
        if (filters.price_ceil) params.set('price_ceil', filters.price_ceil)
        if (filters.page > 1) params.set('page', filters.page)

        setSearchParams(params)
    }

    if (loading) return <div>Загрузка фильтров...</div>
    if (error) return <div>Ошибка загрузки фильтров</div>

    return (
        <div className="filters">
            <div className="filter-section">
                <h3>Категории</h3>
                <div className="filter-options">
                    {categories.map(category => (
                        <label key={category.category_id}>
                            <input
                                type="checkbox"
                                name="category"
                                value={category.category_name}
                                checked={filters.category.includes(category.category_name)}
                                onChange={handleFilterChange}
                            />
                            {category.category_name}
                        </label>
                    ))}
                </div>
            </div>

            <div className="filter-section">
                <h3>Производители</h3>
                <div className="filter-options">
                    {manufacturers.map(manufacturer => (
                        <label key={manufacturer.manufacturer_id}>
                            <input
                                type="checkbox"
                                name="manufacturer"
                                value={manufacturer.manufacturer_name}
                                checked={filters.manufacturer.includes(manufacturer.manufacturer_name)}
                                onChange={handleFilterChange}
                            />
                            {manufacturer.manufacturer_name}
                        </label>
                    ))}
                </div>
            </div>

            <div className="filter-section">
                <h3>Цена</h3>
                <div className="price-inputs">
                    <input
                        type="number"
                        name="price_floor"
                        value={filters.price_floor}
                        onChange={handleFilterChange}
                        placeholder="От"
                    />
                    <input
                        type="number"
                        name="price_ceil"
                        value={filters.price_ceil}
                        onChange={handleFilterChange}
                        placeholder="До"
                    />
                </div>
            </div>
        </div>
    )
} 