import PropTypes from 'prop-types'
import './InstrumentCard.css'

function InstrumentCard({ instrument }) {
    const { name, description, price, image_url, status } = instrument

    return (
        <div className="instrument-card">
            <div className="instrument-card__image">
                <img src={image_url || '/instrument-placeholder.jpg'} alt={name} />
                <div className={`instrument-card__status status-${status}`}>
                    {status === 'available' ? 'Доступен' : 'Арендован'}
                </div>
            </div>
            <div className="instrument-card__content">
                <h3 className="instrument-card__title">{name}</h3>
                <p className="instrument-card__description">{description}</p>
                <div className="instrument-card__price">
                    {price} ₽/день
                </div>
                <button 
                    className="instrument-card__button"
                    disabled={status !== 'available'}
                >
                    {status === 'available' ? 'Арендовать' : 'Недоступен'}
                </button>
            </div>
        </div>
    )
}

InstrumentCard.propTypes = {
    instrument: PropTypes.shape({
        name: PropTypes.string.isRequired,
        description: PropTypes.string.isRequired,
        price: PropTypes.number.isRequired,
        image_url: PropTypes.string,
        status: PropTypes.string.isRequired,
    }).isRequired,
}

export default InstrumentCard 