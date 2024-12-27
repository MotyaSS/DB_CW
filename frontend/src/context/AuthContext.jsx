import { createContext, useContext, useState } from 'react'
import { apiClient } from '../api/client'

const AuthContext = createContext(null)

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null)
    const [token, setToken] = useState(localStorage.getItem('token'))

    const login = async (username, password) => {
        try {
            const { data } = await apiClient.post('/api/sign-in', { username, password })
            localStorage.setItem('token', data.token)
            setToken(data.token)
            // Получаем данные пользователя
            const userResponse = await apiClient.get('/api/users/profile', {
                headers: { Authorization: `Bearer ${data.token}` }
            })
            setUser(userResponse.data)
        } catch (error) {
            console.error('Login error:', error)
            throw error
        }
    }

    const logout = () => {
        localStorage.removeItem('token')
        setToken(null)
        setUser(null)
    }

    const register = async (userData) => {
        try {
            await apiClient.post('/api/sign-up', userData)
        } catch (error) {
            console.error('Registration error:', error)
            throw error
        }
    }

    return (
        <AuthContext.Provider value={{
            user,
            token,
            isAuthenticated: !!token,
            login,
            logout,
            register
        }}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => {
    const context = useContext(AuthContext)
    if (!context) {
        throw new Error('useAuth must be used within AuthProvider')
    }
    return context
} 