import { OTPForm } from "@/components/otp-form"

export default function OTPPage() {
  return (
    <div className="flex min-h-svh w-full">
      <div className="flex w-full items-center justify-center p-6 lg:w-1/2">
        <div className="w-full max-w-xs">
          <OTPForm />
        </div>
      </div>
      <div className="relative hidden w-1/2 lg:block">
        <img
          alt="Authentication"
          height={1080}
          src="/whisper_ico.png"
          width={1920}
          className="object-cover dark:brightness-[0.2] dark:grayscale"
        />
      </div>
    </div>
  )
}
