"use client"

import { useRouter } from "next/navigation"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"

export function ForgotPasswordForm({
  className,
  ...props
}: React.ComponentProps<"form">) {
  const router = useRouter()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // TODO: Enviar email de recuperação aqui
    router.push('/otp')
  }

  return (
    <form onSubmit={handleSubmit} className={cn("flex flex-col gap-4", className)} {...props}>
      <FieldGroup className="gap-3">
        <div className="flex flex-col items-center gap-1 text-center">
          <h1 className="text-2xl font-bold">Reset your password</h1>
          <p className="text-muted-foreground text-sm text-balance">
            Enter your email and we&apos;ll send you a link to reset your password
          </p>
        </div>
        
        <Field>
          <FieldLabel htmlFor="email">Email</FieldLabel>
          <Input 
            id="email" 
            type="email" 
            placeholder="email@example.com" 
            required 
          />
        </Field>

        <Field>
          <Button type="submit" className="w-full">
            Send reset link
          </Button>
        </Field>

        <FieldDescription className="text-center">
          Remember your password?{" "}
          <a href="/login" className="underline underline-offset-4">
            Back to sign in
          </a>
        </FieldDescription>
      </FieldGroup>
    </form>
  )
}
