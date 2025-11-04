"use client"

import { createContext, useContext, useEffect, useState, ReactNode } from "react"
import { 
  User,
  onAuthStateChanged,
  signInWithEmailAndPassword,
  signInWithPopup,
  createUserWithEmailAndPassword,
  signOut as firebaseSignOut,
} from "firebase/auth"
import { auth, googleProvider } from "@/lib/firebase"
import { useRouter } from "next/navigation"

interface AuthContextType {
  user: User | null
  loading: boolean
  login: (email: string, password: string) => Promise<void>
  loginWithGoogle: () => Promise<void>
  register: (email: string, password: string) => Promise<void>
  logout: () => Promise<void>
  getToken: () => Promise<string | null>
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  // Helper para salvar token em cookie E localStorage
  const saveToken = (token: string) => {
    // localStorage (para acesso rápido no client)
    localStorage.setItem("firebase_token", token)
    
    // Cookie (para middleware.ts ter acesso)
    document.cookie = `firebase_token=${token}; path=/; max-age=${60 * 60 * 24 * 7}; SameSite=Lax`
  }

  const clearToken = () => {
    localStorage.removeItem("firebase_token")
    // Remove cookie
    document.cookie = 'firebase_token=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT;'
  }

  // Listener do Firebase - sincroniza estado automaticamente
  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      setUser(user)
      setLoading(false)
      
      // Salva token no localStorage E cookie quando usuário muda
      if (user) {
        user.getIdToken().then(token => {
          saveToken(token)
        })
      } else {
        clearToken()
      }
    })

    return () => unsubscribe()
  }, [])

  // Login com email e senha
  const login = async (email: string, password: string) => {
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password)
      const token = await userCredential.user.getIdToken()
      saveToken(token)
      router.push("/dashboard")
    } catch (error: any) {
      console.error("Login error:", error)
      throw new Error(error.message || "Falha ao fazer login")
    }
  }

  // Login com Google
  const loginWithGoogle = async () => {
    try {
      const userCredential = await signInWithPopup(auth, googleProvider)
      const token = await userCredential.user.getIdToken()
      saveToken(token)
      router.push("/dashboard")
    } catch (error: any) {
      console.error("Google login error:", error)
      throw new Error(error.message || "Falha ao fazer login com Google")
    }
  }

  // Registro de novo usuário
  const register = async (email: string, password: string) => {
    try {
      const userCredential = await createUserWithEmailAndPassword(auth, email, password)
      const token = await userCredential.user.getIdToken()
      saveToken(token)
      router.push("/dashboard")
    } catch (error: any) {
      console.error("Register error:", error)
      throw new Error(error.message || "Falha ao criar conta")
    }
  }

  // Logout
  const logout = async () => {
    try {
      await firebaseSignOut(auth)
      clearToken()
      router.push("/login")
    } catch (error: any) {
      console.error("Logout error:", error)
      throw new Error(error.message || "Falha ao sair")
    }
  }

  // Obter token atualizado
  const getToken = async () => {
    if (!user) return null
    try {
      return await user.getIdToken(true) // força refresh
    } catch (error) {
      console.error("Error getting token:", error)
      return null
    }
  }

  const value = {
    user,
    loading,
    login,
    loginWithGoogle,
    register,
    logout,
    getToken,
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

// Hook para usar o contexto (será exportado em hooks/useAuth.ts)
export function useAuthContext() {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error("useAuthContext must be used within an AuthProvider")
  }
  return context
}
