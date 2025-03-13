import './StoreCard.css'

export default function StoreCard({ store }) {
    return (
        <div className="store-card">
            <div className="store-info">
                <h3>{store.store_name}</h3>
                <div className="store-details">
                    <p className="store-address">
                        <span className="icon">📍</span>
                        {store.store_address}
                    </p>
                    <p className="store-phone">
                        <span className="icon">📞</span>
                        {store.phone_number}
                    </p>
                </div>
            </div>
        </div>
    )
} 