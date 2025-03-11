import { useState } from 'react'
import axios from 'axios'
import CategoryForm from '../components/CategoryForm/CategoryForm'
import ManufacturerForm from '../components/ManufacturerForm/ManufacturerForm'
import Modal from '../components/Modal/Modal'
import InstrumentForm from '../components/InstrumentForm/InstrumentForm'
import './Chief.css'

export default function Chief() {
    const [error, setError] = useState(null)
    const [success, setSuccess] = useState(null)
    const [isCategoryModalOpen, setCategoryModalOpen] = useState(false)
    const [isManufacturerModalOpen, setManufacturerModalOpen] = useState(false)
    const token = localStorage.getItem('token')

    const handleAddCategory = async (categoryData) => {
        try {
            await axios.post('http://localhost:8080/api/instruments/categories', categoryData, {
                headers: { Authorization: `Bearer ${token}` }
            })
            setSuccess('Категория успешно добавлена')
            setError(null)
            setCategoryModalOpen(false)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при добавлении категории')
            setSuccess(null)
        }
    }

    const handleAddManufacturer = async (manufacturerData) => {
        try {
            await axios.post('http://localhost:8080/api/instruments/manufacturers', manufacturerData, {
                headers: { Authorization: `Bearer ${token}` }
            })
            setSuccess('Производитель успешно добавлен')
            setError(null)
            setManufacturerModalOpen(false)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при добавлении производителя')
            setSuccess(null)
        }
    }

    const handleAddInstrument = async (instrumentData) => {
        try {
            await axios.post('http://localhost:8080/api/instruments', instrumentData, {
                headers: { Authorization: `Bearer ${token}` }
            })
        } catch (err) {
            throw err
        }
    }

    return (
        <div className="chief-page">
            <h1>Панель управляющего</h1>
            {error && <div className="error-message">{error}</div>}
            {success && <div className="success-message">{success}</div>}

            <div className="chief-content">
                <section className="management-section">
                    <h2>Управление каталогом</h2>
                    <div className="management-buttons">
                        <button onClick={() => setCategoryModalOpen(true)}>
                            Добавить категорию
                        </button>
                        <button onClick={() => setManufacturerModalOpen(true)}>
                            Добавить производителя
                        </button>
                    </div>
                </section>

                <section className="instruments-section">
                    <h2>Добавить инструмент</h2>
                    <InstrumentForm onSubmit={handleAddInstrument} />
                </section>
            </div>

            <Modal
                isOpen={isCategoryModalOpen}
                onClose={() => setCategoryModalOpen(false)}
                title="Добавить категорию"
            >
                <CategoryForm onSubmit={handleAddCategory} />
            </Modal>

            <Modal
                isOpen={isManufacturerModalOpen}
                onClose={() => setManufacturerModalOpen(false)}
                title="Добавить производителя"
            >
                <ManufacturerForm onSubmit={handleAddManufacturer} />
            </Modal>
        </div>
    )
} 