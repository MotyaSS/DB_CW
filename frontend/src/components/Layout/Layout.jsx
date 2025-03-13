import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { useAuth } from '../../context/AuthContext'
import './Layout.css'

export default function Layout() {
    const navigate = useNavigate()
    const { user, logout } = useAuth()

    const handleAuth = () => {
        if (user) {
            logout()
            navigate('/')
        } else {
            navigate('/login')
        }
    }

    return (
        <div className="layout">
            <header className="layout-header">
                <nav className="layout-nav">
                    <div className="nav-links">
                        <NavLink to="/">Главная</NavLink>
                        <NavLink to="/instruments">Инструменты</NavLink>
                        <NavLink to="/stores">Магазины</NavLink>
                        {user && (
                            <>
                                <NavLink to="/profile">Профиль</NavLink>
                                {user.role_id >= 2 && <NavLink to="/staff">Персонал</NavLink>}
                                {user.role_id >= 3 && <NavLink to="/chief">Менеджер</NavLink>}
                                {user.role_id === 4 && <NavLink to="/admin">Админ</NavLink>}
                            </>
                        )}
                    </div>
                    <div className="auth-section">
                        {user && <span className="username">{user.username}</span>}
                        <button onClick={handleAuth}>
                            {user ? 'Выйти' : 'Войти'}
                        </button>
                    </div>
                </nav>
            </header>
            <main className="layout-main">
                <Outlet />
            </main>
        </div>
    )
} 