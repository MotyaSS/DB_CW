import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AuthProvider } from './context/AuthContext'
import Layout from './components/Layout/Layout'
import Instruments from './pages/Instruments'
import Stores from './pages/Stores'
import Login from './pages/Login'
import Register from './pages/Register'
import Admin from './pages/Admin'
import Staff from './pages/Staff'
import Chief from './pages/Chief'
import Profile from './pages/Profile'
import ProtectedRoute from './components/ProtectedRoute/ProtectedRoute'
import homepageImage from './assets/images/homepage.png'
import InstrumentDetails from './pages/InstrumentDetails'
import './App.css'

function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="instruments" element={<Instruments />} />
            <Route path="instruments/:id" element={<InstrumentDetails />} />
            <Route path="stores" element={<Stores />} />
            <Route path="login" element={<Login />} />
            <Route path="register" element={<Register />} />
            <Route path="profile" element={
              <ProtectedRoute>
                <Profile />
              </ProtectedRoute>
            } />
            <Route path="staff" element={
              <ProtectedRoute requiredRole={2}>
                <Staff />
              </ProtectedRoute>
            } />
            <Route path="chief" element={
              <ProtectedRoute requiredRole={3}>
                <Chief />
              </ProtectedRoute>
            } />
            <Route path="admin" element={
              <ProtectedRoute requiredRole={4}>
                <Admin />
              </ProtectedRoute>
            } />
          </Route>
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  )
}

function Home() {
  return (
    <div className="home" style={{
      backgroundImage: `url(${homepageImage})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
      position: 'relative'
    }}>
      <div className="overlay"></div>
      <h1>Прокат музыкальных инструментов</h1>
      <p>Найдите идеальный инструмент для вашего творчества</p>
      <div className="home-actions">
        <a href="/instruments" className="cta-button">
          Смотреть инструменты
        </a>
      </div>
    </div>
  )
}

export default App
