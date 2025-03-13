import { useEffect, useState } from 'react'
import { useAuth } from '../context/AuthContext'
import CreateUserForm from '../components/CreateUserForm/CreateUserForm'
import './AdminPage.css'

export default function AdminPage() {
    const { user } = useAuth()

    return (
        <div className="admin-page">
            <h1>Панель администратора</h1>
            <CreateUserForm
                roleId={2} // ID роли для менеджера
                title="Создать аккаунт менеджера"
            />
            {/* Остальной контент страницы администратора */}
        </div>
    )
} 