import { createContext, useContext, useState, useEffect } from 'react'
import axios from 'axios'

const AuthContext = createContext(null)

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        checkAuth()
    }, [])

    const checkAuth = async () => {
        const token = localStorage.getItem('token')
        if (token) {
            try {
                axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
                const response = await axios.get('http://localhost:8080/api/auth/me')
                setUser(response.data)
            } catch (err) {
                console.error('Auth check failed:', err)
                localStorage.removeItem('token')
                delete axios.defaults.headers.common['Authorization']
            }
        }
        setLoading(false)
    }

    const login = async (username, password) => {
        const response = await axios.post('http://localhost:8080/api/auth/sign-in', {
            username,
            password
        })
        const { token } = response.data
        localStorage.setItem('token', token)
        axios.defaults.headers.common['Authorization'] = `Bearer ${token}`

        const userResponse = await axios.get('http://localhost:8080/api/auth/me')
        setUser(userResponse.data)
        return userResponse.data
    }

    const register = async (userData) => {
        await axios.post('http://localhost:8080/api/auth/sign-up', userData)
    }

    const logout = () => {
        localStorage.removeItem('token')
        delete axios.defaults.headers.common['Authorization']
        setUser(null)
    }

    const value = {
        user,
        loading,
        login,
        logout,
        register,
        checkAuth
    }

    return (
        <AuthContext.Provider value={value}>
            {!loading && children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    return useContext(AuthContext)
} 