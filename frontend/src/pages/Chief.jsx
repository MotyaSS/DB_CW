import { useState, useEffect } from 'react'
import { useAuth } from '../context/AuthContext'
import axios from 'axios'
import CategoryForm from '../components/CategoryForm/CategoryForm'
import ManufacturerForm from '../components/ManufacturerForm/ManufacturerForm'
import StoreForm from '../components/StoreForm/StoreForm'
import Modal from '../components/Modal/Modal'
import InstrumentForm from '../components/InstrumentForm/InstrumentForm'
import StoreDeleteForm from '../components/StoreDeleteForm/StoreDeleteForm'
import CreateUserForm from '../components/CreateUserForm/CreateUserForm'
import './Chief.css'

export default function Chief() {
    const { user } = useAuth()
    const [instruments, setInstruments] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)
    const [success, setSuccess] = useState(null)
    const [isCategoryModalOpen, setCategoryModalOpen] = useState(false)
    const [isManufacturerModalOpen, setManufacturerModalOpen] = useState(false)
    const [isStoreModalOpen, setStoreModalOpen] = useState(false)
    const [isStoreDeleteModalOpen, setStoreDeleteModalOpen] = useState(false)
    const [isInstrumentModalOpen, setInstrumentModalOpen] = useState(false)
    const [isCreateUserModalOpen, setCreateUserModalOpen] = useState(false)
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
            setSuccess('Инструмент успешно добавлен')
            setError(null)
            setInstrumentModalOpen(false)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при добавлении инструмента')
            setSuccess(null)
        }
    }

    const handleAddStore = async (storeData) => {
        try {
            await axios.post('http://localhost:8080/api/stores', storeData, {
                headers: { Authorization: `Bearer ${token}` }
            })
            setSuccess('Магазин успешно добавлен')
            setError(null)
            setStoreModalOpen(false)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при добавлении магазина')
            setSuccess(null)
        }
    }

    const handleDeleteStore = async (storeId) => {
        try {
            await axios.delete(`http://localhost:8080/api/stores/${storeId}`, {
                headers: { Authorization: `Bearer ${token}` }
            })
            setSuccess('Магазин успешно удален')
            setError(null)
            setStoreDeleteModalOpen(false)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при удалении магазина')
            setSuccess(null)
        }
    }

    return (
        <div className="chief-page">
            <div className="chief-header">
                <h1>Панель менеджера</h1>
                <button
                    className="create-user-btn"
                    onClick={() => setCreateUserModalOpen(true)}
                >
                    Создать аккаунт персонала
                </button>
            </div>
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

                <section className="management-card">
                    <h2>Управление магазинами</h2>
                    <div className="management-buttons">
                        <button onClick={() => setStoreModalOpen(true)}>
                            Добавить магазин
                        </button>
                        <button
                            onClick={() => setStoreDeleteModalOpen(true)}
                            className="delete-button"
                        >
                            Удалить магазин
                        </button>
                    </div>
                </section>

                <section className="management-card">
                    <h2>Управление инструментами</h2>
                    <div className="management-buttons">
                        <button onClick={() => setInstrumentModalOpen(true)}>
                            Добавить инструмент
                        </button>
                    </div>
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

            <Modal
                isOpen={isStoreModalOpen}
                onClose={() => setStoreModalOpen(false)}
                title="Добавить магазин"
            >
                <StoreForm onSubmit={handleAddStore} />
            </Modal>

            <Modal
                isOpen={isStoreDeleteModalOpen}
                onClose={() => setStoreDeleteModalOpen(false)}
                title="Удалить магазин"
            >
                <StoreDeleteForm onSubmit={handleDeleteStore} />
            </Modal>

            <Modal
                isOpen={isInstrumentModalOpen}
                onClose={() => setInstrumentModalOpen(false)}
                title="Добавить инструмент"
            >
                <InstrumentForm onSubmit={handleAddInstrument} />
            </Modal>

            <Modal
                isOpen={isCreateUserModalOpen}
                onClose={() => setCreateUserModalOpen(false)}
                title="Создать аккаунт персонала"
            >
                <CreateUserForm
                    roleId={3}
                    title="Создать аккаунт персонала"
                />
            </Modal>
        </div>
    )
} 