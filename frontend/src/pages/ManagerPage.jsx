import { useEffect, useState } from 'react'
import { useAuth } from '../context/AuthContext'
import CreateUserForm from '../components/CreateUserForm/CreateUserForm'
import './ManagerPage.css'

export default function ManagerPage() {
    const { user } = useAuth()

    return (
        <div className="manager-page">
            <h1>Панель менеджера</h1>
            <CreateUserForm
                roleId={3} // ID роли для персонала
                title="Создать аккаунт персонала"
            />
            {/* Остальной контент страницы менеджера */}
        </div>
    )
} 