import { useAuthContext } from "@/context/AuthContext"

/**
 * Hook customizado para acessar o contexto de autenticação
 * 
 * Uso:
 * const { user, loading, login, logout } = useAuth()
 * 
 * @returns {Object} - Objeto com dados e métodos de autenticação
 */
export function useAuth() {
  return useAuthContext()
}
