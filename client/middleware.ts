import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

/**
 * Middleware do Next.js para proteção de rotas
 * Roda no Edge Runtime (servidor) ANTES de carregar a página
 * 
 * Funciona como primeira camada de defesa
 */
export function middleware(request: NextRequest) {
  const token = request.cookies.get('firebase_token')?.value
  const { pathname } = request.nextUrl

  // Lista de rotas que precisam de autenticação
  const protectedPaths = [
    '/dashboard',
    '/profile',
    '/settings',
    '/messages',
    '/admin',
  ]

  // Verifica se o path atual é protegido
  const isProtectedPath = protectedPaths.some(path => 
    pathname.startsWith(path)
  )

  // Se é rota protegida e não tem token, redireciona
  if (isProtectedPath && !token) {
    const loginUrl = new URL('/login', request.url)
    // Adiciona redirect URL para voltar depois do login
    loginUrl.searchParams.set('redirect', pathname)
    return NextResponse.redirect(loginUrl)
  }

  // Se está na página de login/register e JÁ tem token, redireciona para dashboard
  const publicAuthPaths = ['/login', '/register']
  const isPublicAuthPath = publicAuthPaths.some(path => pathname.startsWith(path))
  
  if (isPublicAuthPath && token) {
    return NextResponse.redirect(new URL('/dashboard', request.url))
  }

  // TODO (Opcional): Verificar validade do token no Firebase Admin SDK
  // Isso requer Firebase Admin SDK no servidor e é mais complexo
  // try {
  //   const admin = await import('firebase-admin')
  //   const decodedToken = await admin.auth().verifyIdToken(token)
  //   // Token válido, continua
  // } catch (error) {
  //   // Token inválido, redireciona
  //   return NextResponse.redirect(new URL('/login', request.url))
  // }

  return NextResponse.next()
}

// Configura quais rotas o middleware deve interceptar
export const config = {
  matcher: [
    /*
     * Match all request paths except for:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico, sitemap.xml, robots.txt (static files)
     */
    '/((?!api|_next/static|_next/image|favicon.ico|.*\\.svg|.*\\.png|.*\\.jpg).*)',
  ],
}

