import { useState } from 'react'
import './CategoryForm.css'

export default function CategoryForm({ onSubmit }) {
    const [formData, setFormData] = useState({
        category_name: '',
        category_description: ''
    })

    const handleSubmit = (e) => {
        e.preventDefault()
        onSubmit(formData)
        setFormData({ category_name: '', category_description: '' })
    }

    return (
        <form className="category-form" onSubmit={handleSubmit}>
            <div className="form-group">
                <label htmlFor="category_name">Название категории</label>
                <input
                    type="text"
                    id="category_name"
                    value={formData.category_name}
                    onChange={(e) => setFormData(prev => ({ ...prev, category_name: e.target.value }))}
                    required
                />
            </div>
            <div className="form-group">
                <label htmlFor="category_description">Описание</label>
                <textarea
                    id="category_description"
                    value={formData.category_description}
                    onChange={(e) => setFormData(prev => ({ ...prev, category_description: e.target.value }))}
                />
            </div>
            <button type="submit">Добавить категорию</button>
        </form>
    )
} 