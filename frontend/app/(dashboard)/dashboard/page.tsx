'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { useAuthStore } from '@/store/auth'
import { propertyApi, bookingApi } from '@/lib/api'
import { Property, Booking } from '@/types'

export default function DashboardPage() {
  const { user, isAuthenticated } = useAuthStore()
  const router = useRouter()
  const [properties, setProperties] = useState<Property[]>([])
  const [bookings, setBookings] = useState<Booking[]>([])
  const [loading, setLoading] = useState(true)

  // 安全获取数组长度的辅助函数
  const safeLength = (arr: any[] | null | undefined): number => {
    return Array.isArray(arr) ? arr.length : 0
  }

  // 安全过滤数组的辅助函数
  const safeFilter = (arr: any[] | null | undefined, filterFn: (item: any) => boolean): any[] => {
    return Array.isArray(arr) ? arr.filter(filterFn) : []
  }

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/login')
      return
    }

    const fetchData = async () => {
      try {
        if (user?.role === 'landlord') {
          // Fetch properties for landlords
          try {
            const propertiesData = await propertyApi.getProperties()
            setProperties(Array.isArray(propertiesData) ? propertiesData.slice(0, 5) : []) // Show latest 5
          } catch (error) {
            console.error('Failed to fetch properties:', error)
            setProperties([])
          }
        }
        
        // Fetch bookings for all users
        try {
          const bookingsData = await bookingApi.getBookings({ limit: 5 })
          setBookings(Array.isArray(bookingsData) ? bookingsData : [])
        } catch (error) {
          console.error('Failed to fetch bookings:', error)
          setBookings([])
        }
      } catch (error) {
        console.error('Failed to fetch dashboard data:', error)
        setProperties([])
        setBookings([])
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [isAuthenticated, user, router])

  if (!isAuthenticated) {
    return null
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-lg">加载中...</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Navigation */}
      <nav className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <h1 className="text-xl font-semibold">RentHelp Dashboard</h1>
            </div>
            <div className="flex items-center space-x-4">
              <span className="text-gray-700">欢迎, {user?.name}</span>
              <Button variant="outline" onClick={() => router.push('/properties')}>
                浏览房源
              </Button>
              {user?.role === 'landlord' && (
                <Button onClick={() => router.push('/properties/new')}>
                  发布房源
                </Button>
              )}
              <Button variant="ghost" onClick={() => useAuthStore.getState().logout()}>
                退出
              </Button>
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Welcome Section */}
        <div className="px-4 py-6 sm:px-0">
          <Card>
            <CardHeader>
              <CardTitle>欢迎来到 RentHelp</CardTitle>
              <CardDescription>
                {user?.role === 'landlord' 
                  ? '管理您的房源，查看预订情况'
                  : '找到您的理想住所，管理您的预订'
                }
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="bg-blue-50 p-4 rounded-lg">
                  <h3 className="font-semibold text-blue-900">
                    {user?.role === 'landlord' ? '我的房源' : '我的预订'}
                  </h3>
                  <p className="text-2xl font-bold text-blue-700">
                    {user?.role === 'landlord' ? safeLength(properties) : safeLength(bookings)}
                  </p>
                </div>
                <div className="bg-green-50 p-4 rounded-lg">
                  <h3 className="font-semibold text-green-900">活跃预订</h3>
                  <p className="text-2xl font-bold text-green-700">
                    {safeLength(safeFilter(bookings, (b: any) => b.status === 'confirmed'))}
                  </p>
                </div>
                <div className="bg-yellow-50 p-4 rounded-lg">
                  <h3 className="font-semibold text-yellow-900">待处理</h3>
                  <p className="text-2xl font-bold text-yellow-700">
                    {safeLength(safeFilter(bookings, (b: any) => b.status === 'pending'))}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Properties Section (for landlords) */}
        {user?.role === 'landlord' && (
          <div className="px-4 py-6 sm:px-0">
            <Card>
              <CardHeader>
                <CardTitle>最新房源</CardTitle>
                <CardDescription>您最近发布的房源</CardDescription>
              </CardHeader>
              <CardContent>
                {safeLength(properties) === 0 ? (
                  <div className="text-center py-8">
                    <p className="text-gray-500">您还没有发布任何房源</p>
                    <Button className="mt-4" onClick={() => router.push('/properties/new')}>
                      发布第一个房源
                    </Button>
                  </div>
                ) : (
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {properties?.map((property) => (
                      <div key={property?.id || property?._id} className="border rounded-lg p-4">
                        <h4 className="font-semibold">{property?.title}</h4>
                        <p className="text-sm text-gray-600 mt-1">{property?.location?.city}</p>
                        <p className="text-lg font-bold mt-2">¥{property?.price}/月</p>
                        <div className="flex justify-between items-center mt-3">
                          <span className={`px-2 py-1 rounded text-xs ${
                            property.available 
                              ? 'bg-green-100 text-green-800' 
                              : 'bg-red-100 text-red-800'
                          }`}>
                            {property.available ? '可租' : '已租'}
                          </span>
                          <Button size="sm" variant="outline">
                            查看详情
                          </Button>
                        </div>
                      </div>
                    )) || []}
                  </div>
                )}
              </CardContent>
            </Card>
          </div>
        )}

        {/* Bookings Section */}
        <div className="px-4 py-6 sm:px-0">
          <Card>
            <CardHeader>
              <CardTitle>
                {user?.role === 'landlord' ? '房源预订' : '我的预订'}
              </CardTitle>
              <CardDescription>
                {user?.role === 'landlord' 
                  ? '租客对您房源的预订申请'
                  : '您的租房预订记录'
                }
              </CardDescription>
            </CardHeader>
            <CardContent>
              {safeLength(bookings) === 0 ? (
                <div className="text-center py-8">
                  <p className="text-gray-500">
                    {user?.role === 'landlord' 
                      ? '暂无预订申请'
                      : '您还没有任何预订记录'
                    }
                  </p>
                </div>
              ) : (
                <div className="space-y-4">
                  {bookings?.map((booking) => (
                    <div key={booking?.id || booking?._id} className="border rounded-lg p-4">
                      <div className="flex justify-between items-start">
                        <div>
                          <h4 className="font-semibold">预订 #{booking.id?.slice(-6) || 'N/A'}</h4>
                          <p className="text-sm text-gray-600 mt-1">
                            {new Date(booking.start_date).toLocaleDateString()} - {new Date(booking.end_date).toLocaleDateString()}
                          </p>
                          <p className="text-lg font-bold mt-2">¥{booking.total_price || 0}</p>
                        </div>
                        <span className={`px-2 py-1 rounded text-xs ${
                          booking.status === 'confirmed' ? 'bg-green-100 text-green-800' :
                          booking.status === 'pending' ? 'bg-yellow-100 text-yellow-800' :
                          booking.status === 'cancelled' ? 'bg-red-100 text-red-800' :
                          'bg-gray-100 text-gray-800'
                        }`}>
                          {booking.status === 'confirmed' ? '已确认' :
                           booking.status === 'pending' ? '待确认' :
                           booking.status === 'cancelled' ? '已取消' : '已完成'}
                        </span>
                      </div>
                    </div>
                  )) || []}
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      </main>
    </div>
  )
}
