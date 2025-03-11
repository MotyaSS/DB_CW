import { useState } from 'react'
import './UserList.css'

export default function UserList({ users, onDelete }) {
    const getRoleName = (roleId) => {
        const roles = {
            1: 'Customer',
            2: 'Staff',
            3: 'Chief',
            4: 'Admin'
        }
        return roles[roleId] || 'Unknown'
    }

    return (
        <div className="user-list">
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Имя пользователя</th>
                        <th>Email</th>
                        <th>Телефон</th>
                        <th>Роль</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
                    {users.map(user => (
                        <tr key={user.user_id}>
                            <td>{user.user_id}</td>
                            <td>{user.username}</td>
                            <td>{user.email}</td>
                            <td>{user.phone_number}</td>
                            <td>{getRoleName(user.role_id)}</td>
                            <td>
                                <button
                                    className="delete-btn"
                                    onClick={() => onDelete(user.user_id)}
                                >
                                    Удалить
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
} 