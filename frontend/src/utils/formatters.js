export const formatPrice = (price) => {
    if (!price) return '0'
    // Преобразуем decimal в число и форматируем с двумя знаками после запятой
    return Number(price).toFixed(2)
} 