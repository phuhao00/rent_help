import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { useAuthStore } from '@/store/auth'

export const useAuthRedirect = (redirectTo = '/login') => {
  const { isAuthenticated } = useAuthStore()
  const router = useRouter()

  useEffect(() => {
    if (!isAuthenticated) {
      router.push(redirectTo)
    }
  }, [isAuthenticated, redirectTo, router])

  return isAuthenticated
}

export const useRequireAuth = () => {
  const { isAuthenticated, user } = useAuthStore()
  const router = useRouter()

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
    }
  }, [isAuthenticated, router])

  return { isAuthenticated, user }
}

export const useRoleRedirect = (allowedRoles: string[], redirectTo = '/dashboard') => {
  const { user } = useAuthStore()
  const router = useRouter()

  useEffect(() => {
    if (user && !allowedRoles.includes(user.role)) {
      router.push(redirectTo)
    }
  }, [user, allowedRoles, redirectTo, router])

  return user?.role
}
