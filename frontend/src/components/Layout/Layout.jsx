import { Outlet, NavLink } from 'react-router-dom'
import { useAuth } from '../../context/AuthContext'
import './Layout.css'

function Layout() {
    const { isAuthenticated, user, logout } = useAuth()

    return (
        <div className="layout">
            <header className="layout-header">
                <nav className="layout-nav">
                    <NavLink to="/">Главная</NavLink>
                    <NavLink to="/instruments">Инструменты</NavLink>
                    <NavLink to="/stores">Магазины</NavLink>
                    {isAuthenticated ? (
                        <>
                            <NavLink to="/profile">Профиль</NavLink>
                            {user?.roleId >= 2 && <NavLink to="/admin">Админ панель</NavLink>}
                            <button onClick={logout}>Выйти</button>
                        </>
                    ) : (
                        <NavLink to="/auth">Войти</NavLink>
                    )}
                </nav>
            </header>
            <main className="layout-main">
                <Outlet />
            </main>
        </div>
    )
}

export default Layout 