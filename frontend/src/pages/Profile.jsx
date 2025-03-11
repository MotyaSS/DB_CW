import { useAuth } from '../context/AuthContext'
import { getRoleName } from '../utils/roleUtils'
import './Profile.css'

export default function Profile() {
    const { user } = useAuth()

    return (
        <div className="profile-page">
            <h1>Профиль</h1>
            <div className="profile-content">
                <div className="profile-info">
                    <h2>Личная информация</h2>
                    <div className="info-group">
                        <label>ID:</label>
                        <span>{user.user_id}</span>
                    </div>
                    <div className="info-group">
                        <label>Имя пользователя:</label>
                        <span>{user.username}</span>
                    </div>
                    <div className="info-group">
                        <label>Email:</label>
                        <span>{user.email}</span>
                    </div>
                    <div className="info-group">
                        <label>Телефон:</label>
                        <span>{user.phone_number}</span>
                    </div>
                    <div className="info-group">
                        <label>Роль:</label>
                        <span className="role-badge">{getRoleName(user.role_id)}</span>
                    </div>
                </div>
            </div>
        </div>
    )
} 