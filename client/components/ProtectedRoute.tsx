"use client"

import { useEffect } from "react"
import { useRouter } from "next/navigation"
import { useAuth } from "@/hooks/useAuth"

interface ProtectedRouteProps {
  children: React.ReactNode
  redirectTo?: string
}

/**
 * Componente para proteger rotas que requerem autenticação
 * 
 * Uso:
 * <ProtectedRoute>
 *   <DashboardContent />
 * </ProtectedRoute>
 */
export function ProtectedRoute({ 
  children, 
  redirectTo = "/login" 
}: ProtectedRouteProps) {
  const { user, loading } = useAuth()
  const router = useRouter()

  useEffect(() => {
    if (!loading && !user) {
      router.push(redirectTo)
    }
  }, [user, loading, router, redirectTo])

  // Mostra loading enquanto verifica autenticação
  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="flex flex-col items-center gap-4">
          <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
          <p className="text-sm text-muted-foreground">Loading...</p>
        </div>
      </div>
    )
  }

  // Se não está logado, não renderiza nada (vai redirecionar)
  if (!user) {
    return null
  }

  // Usuário autenticado, renderiza o conteúdo
  return <>{children}</>
}

