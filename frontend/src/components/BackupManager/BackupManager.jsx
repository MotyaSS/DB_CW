import { useState, useEffect } from 'react'
import axios from 'axios'
import './BackupManager.css'

export default function BackupManager() {
    const [backups, setBackups] = useState([])
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)
    const [status, setStatus] = useState(null)

    const fetchBackups = async () => {
        try {
            setLoading(true)
            const response = await axios.get('http://localhost:8080/api/backups', {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            })
            setBackups(response.data.backups)
        } catch (err) {
            setError(err.response?.data?.message || 'Ошибка при загрузке списка бэкапов')
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchBackups()
    }, [])

    const createBackup = async () => {
        try {
            setLoading(true)
            const response = await axios.post('http://localhost:8080/api/backups', {}, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`
                }
            })
            setStatus({ type: 'success', message: 'Бэкап успешно создан' })
            fetchBackups()
        } catch (err) {
            setStatus({ type: 'error', message: err.response?.data?.message || 'Ошибка при создании бэкапа' })
        } finally {
            setLoading(false)
        }
    }

    const restoreBackup = async (backupPath) => {
        if (!window.confirm('Вы уверены, что хотите восстановить базу данных из этого бэкапа? Все текущие данные будут заменены.')) {
            return
        }

        try {
            setLoading(true)
            await axios.post('http://localhost:8080/api/backups/restore',
                { backup_path: backupPath },
                {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem('token')}`
                    }
                }
            )
            setStatus({ type: 'success', message: 'База данных успешно восстановлена из бэкапа' })
        } catch (err) {
            setStatus({ type: 'error', message: err.response?.data?.message || 'Ошибка при восстановлении из бэкапа' })
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="backup-manager">
            <div className="backup-header">
                <h2>Управление резервными копиями</h2>
                <button
                    onClick={createBackup}
                    disabled={loading}
                    className="create-backup-btn"
                >
                    Создать бэкап
                </button>
            </div>

            {status && (
                <div className={`status-message ${status.type}`}>
                    {status.message}
                </div>
            )}

            {loading && <div className="loading">Загрузка...</div>}

            {error && <div className="error-message">{error}</div>}

            <div className="backups-list">
                <h3>Список доступных бэкапов:</h3>
                {backups.length === 0 ? (
                    <p className="no-backups">Нет доступных бэкапов</p>
                ) : (
                    <ul>
                        {backups.map((backup, index) => (
                            <li key={index} className="backup-item">
                                <span className="backup-name">{backup}</span>
                                <button
                                    onClick={() => restoreBackup(backup)}
                                    disabled={loading}
                                    className="restore-btn"
                                >
                                    Восстановить
                                </button>
                            </li>
                        ))}
                    </ul>
                )}
            </div>
        </div>
    )
} 