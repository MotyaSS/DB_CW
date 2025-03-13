import './Review.css'

export default function Review({ review }) {
    const formatDate = (dateString) => {
        return new Date(dateString).toLocaleDateString('ru-RU', {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        })
    }

    return (
        <div className="review">
            <div className="review-header">
                <div className="review-rating">
                    {'★'.repeat(review.rating)}
                    {'☆'.repeat(5 - review.rating)}
                </div>
                <div className="review-date">{formatDate(review.created_at)}</div>
            </div>
            <div className="review-text">{review.review_text}</div>
            <div className="review-author">
                {review.user_name || 'Анонимный пользователь'}
            </div>
        </div>
    )
}
