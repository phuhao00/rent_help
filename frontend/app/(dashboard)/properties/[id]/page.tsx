'use client'

import { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { propertyApi, bookingApi } from '@/lib/api'
import { useAuthStore } from '@/store/auth'
import { Property } from '@/types'

export default function PropertyDetailPage() {
  const params = useParams()
  const router = useRouter()
  const { user } = useAuthStore()
  const [property, setProperty] = useState<Property | null>(null)
  const [loading, setLoading] = useState(true)
  const [bookingLoading, setBookingLoading] = useState(false)

  useEffect(() => {
    if (params.id) {
      fetchProperty(params.id as string)
    }
  }, [params.id])

  const fetchProperty = async (id: string) => {
    try {
      const data = await propertyApi.getProperty(id)
      setProperty(data)
    } catch (error) {
      console.error('Failed to fetch property:', error)
      router.push('/properties')
    } finally {
      setLoading(false)
    }
  }

  const handleBooking = async () => {
    if (!property || !user) return

    setBookingLoading(true)
    try {
      // 简单的预订逻辑，实际应用中应该有日期选择等
      const startDate = new Date()
      const endDate = new Date()
      endDate.setMonth(endDate.getMonth() + 1) // 默认租一个月

      await bookingApi.createBooking({
        property_id: property.id,
        start_date: startDate.toISOString(),
        end_date: endDate.toISOString(),
        total_price: property.price,
        message: '我对这个房源感兴趣，希望预订'
      })

      alert('预订申请已提交！')
      router.push('/dashboard')
    } catch (error) {
      console.error('Failed to create booking:', error)
      alert('预订失败，请重试')
    } finally {
      setBookingLoading(false)
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-lg">加载中...</div>
      </div>
    )
  }

  if (!property) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-lg text-gray-600">房源不存在</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Back Button */}
        <div className="mb-6">
          <Button variant="outline" onClick={() => router.back()}>
            ← 返回
          </Button>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Main Content */}
          <div className="lg:col-span-2">
            {/* Images */}
            <div className="mb-6">
              {property.images && property.images.length > 0 ? (
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  {property.images.slice(0, 4).map((image, index) => (
                    <div key={index} className="aspect-w-16 aspect-h-9">
                      <img
                        src={image}
                        alt={`${property.title} - 图片 ${index + 1}`}
                        className="w-full h-64 object-cover rounded-lg"
                      />
                    </div>
                  ))}
                </div>
              ) : (
                <div className="w-full h-64 bg-gray-200 rounded-lg flex items-center justify-center">
                  <span className="text-gray-400">暂无图片</span>
                </div>
              )}
            </div>

            {/* Property Info */}
            <Card>
              <CardHeader>
                <div className="flex justify-between items-start">
                  <div>
                    <CardTitle className="text-2xl">{property.title}</CardTitle>
                    <CardDescription className="text-lg mt-2">
                      {property.location.address}, {property.location.city}
                    </CardDescription>
                  </div>
                  <Badge variant={property.available ? "default" : "destructive"}>
                    {property.available ? '可租' : '已租'}
                  </Badge>
                </div>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div>
                    <h3 className="font-semibold mb-2">房源描述</h3>
                    <p className="text-gray-600">{property.description}</p>
                  </div>

                  <div>
                    <h3 className="font-semibold mb-2">房源信息</h3>
                    <div className="grid grid-cols-2 gap-4">
                      <div>
                        <span className="text-gray-500">类型：</span>
                        <span className="ml-2">
                          {property.type === 'apartment' ? '公寓' :
                           property.type === 'house' ? '别墅' :
                           property.type === 'room' ? '单间' : property.type}
                        </span>
                      </div>
                      <div>
                        <span className="text-gray-500">价格：</span>
                        <span className="ml-2 font-semibold text-lg">
                          ¥{property.price} / 月
                        </span>
                      </div>
                    </div>
                  </div>

                  {property.features && property.features.length > 0 && (
                    <div>
                      <h3 className="font-semibold mb-2">房源特色</h3>
                      <div className="flex flex-wrap gap-2">
                        {property.features.map((feature, index) => (
                          <Badge key={index} variant="secondary">
                            {feature}
                          </Badge>
                        ))}
                      </div>
                    </div>
                  )}

                  <div>
                    <h3 className="font-semibold mb-2">位置信息</h3>
                    <div className="space-y-2">
                      <p><span className="text-gray-500">完整地址：</span>{property.location.address}</p>
                      <p><span className="text-gray-500">城市：</span>{property.location.city}</p>
                      <p><span className="text-gray-500">省份：</span>{property.location.state}</p>
                      <p><span className="text-gray-500">国家：</span>{property.location.country}</p>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>

          {/* Sidebar */}
          <div className="lg:col-span-1">
            <Card className="sticky top-8">
              <CardHeader>
                <CardTitle>预订信息</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="text-center">
                    <div className="text-3xl font-bold text-blue-600">
                      ¥{property.price}
                    </div>
                    <div className="text-gray-500">每月</div>
                  </div>

                  {property.available ? (
                    user?.role === 'tenant' ? (
                      user?.id !== property.owner_id ? (
                        <Button 
                          className="w-full" 
                          onClick={handleBooking}
                          disabled={bookingLoading}
                        >
                          {bookingLoading ? '提交中...' : '立即预订'}
                        </Button>
                      ) : (
                        <div className="text-center text-gray-500">
                          这是您发布的房源
                        </div>
                      )
                    ) : (
                      <div className="text-center text-gray-500">
                        只有租客可以预订房源
                      </div>
                    )
                  ) : (
                    <Button className="w-full" disabled>
                      房源已被预订
                    </Button>
                  )}

                  <div className="text-sm text-gray-500 text-center">
                    预订后房东会尽快联系您
                  </div>
                </div>
              </CardContent>
            </Card>

            {/* Contact Info */}
            <Card className="mt-6">
              <CardHeader>
                <CardTitle>联系方式</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-2">
                  <p className="text-sm text-gray-600">
                    如有任何问题，请通过平台消息联系房东
                  </p>
                  <Button variant="outline" className="w-full">
                    发送消息
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  )
}
