import { Link } from 'react-router-dom'
import { formatPrice } from '../../utils/formatters'
import './InstrumentCard.css'

export default function InstrumentCard({ instrument }) {
    const defaultImage = 'https://placehold.co/600x400/1a1f2e/646cff?text=Нет+изображения'

    // Получаем данные из вложенного объекта Instrument
    const {
        instrument_id,
        instrument_name,
        description,
        price_per_day,
        image_url
    } = instrument.Instrument || instrument // Поддерживаем оба варианта структуры

    // Добавляем логи
    console.log('Raw instrument:', instrument)
    console.log('Extracted image_url:', image_url)
    console.log('Final image source:', image_url || defaultImage)

    return (
        <Link to={`/instruments/${instrument_id}`} className="instrument-card">
            <div className="instrument-image">
                <img
                    src={image_url || defaultImage}
                    alt={instrument_name}
                    onError={(e) => {
                        console.log('Image load error for:', instrument_name)
                        e.target.src = defaultImage
                    }}
                />
            </div>
            <div className="instrument-info">
                <h3>{instrument_name}</h3>
                <p className="description">{description}</p>
                <p className="price">{formatPrice(price_per_day)} ₽/день</p>
                {instrument.Discount && (
                    <div className="discount">
                        Скидка: {formatPrice(instrument.Discount.discount_percentage)}%
                    </div>
                )}
            </div>
        </Link>
    )
}
