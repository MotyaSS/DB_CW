import './StoreCard.css'

export default function StoreCard({ store }) {
    return (
        <div className="store-card">
            <div className="store-info">
                <h3>{store.name}</h3>
                <div className="store-details">
                    <p className="store-title">{store.title || "Музыкальный магазин"}</p>
                    <p className="store-address">
                        <span className="icon">📍</span>
                        {store.address}
                    </p>
                    <p className="store-phone">
                        <span className="icon">📞</span>
                        {store.phone}
                    </p>
                    <p className="store-hours">
                        <span className="icon">🕒</span>
                        Часы работы: {store.opening_hours || "Не указаны"}
                    </p>
                </div>
            </div>
        </div>
    )
} 