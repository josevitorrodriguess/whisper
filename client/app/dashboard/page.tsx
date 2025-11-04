"use client"

import { ProtectedRoute } from "@/components/ProtectedRoute"
import { useAuth } from "@/hooks/useAuth"
import { Button } from "@/components/ui/button"

export default function DashboardPage() {
  return (
    <ProtectedRoute>
      <DashboardContent />
    </ProtectedRoute>
  )
}

function DashboardContent() {
  const { user, logout } = useAuth()

  return (
    <div className="flex min-h-screen flex-col items-center justify-center gap-8 p-8">
      <div className="flex flex-col items-center gap-4 text-center">
        <h1 className="text-4xl font-bold">Dashboard</h1>
        <p className="text-muted-foreground">
          Welcome to your protected dashboard!
        </p>
      </div>

      <div className="flex flex-col gap-4 rounded-lg border p-6 min-w-[400px]">
        <h2 className="text-xl font-semibold">User Information</h2>
        
        <div className="space-y-2">
          <div>
            <span className="text-sm text-muted-foreground">Email:</span>
            <p className="font-medium">{user?.email || "N/A"}</p>
          </div>
          
          <div>
            <span className="text-sm text-muted-foreground">User ID:</span>
            <p className="font-mono text-sm">{user?.uid || "N/A"}</p>
          </div>
          
          <div>
            <span className="text-sm text-muted-foreground">Display Name:</span>
            <p className="font-medium">{user?.displayName || "Not set"}</p>
          </div>
          
          <div>
            <span className="text-sm text-muted-foreground">Email Verified:</span>
            <p className="font-medium">
              {user?.emailVerified ? "✅ Yes" : "❌ No"}
            </p>
          </div>
        </div>

        <Button onClick={logout} variant="destructive" className="mt-4">
          Logout
        </Button>
      </div>
    </div>
  )
}

