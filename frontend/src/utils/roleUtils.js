export const getRoleName = (roleId) => {
    const roles = {
        1: 'Покупатель',
        2: 'Сотрудник',
        3: 'Управляющий',
        4: 'Администратор'
    }
    return roles[roleId] || 'Неизвестная роль'
} 