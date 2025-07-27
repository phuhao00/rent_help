'use client'

import { useEffect, useState } from 'react'
import Link from 'next/link'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { propertyApi } from '@/lib/api'
import { Property } from '@/types'

export default function PropertiesPage() {
  const [properties, setProperties] = useState<Property[]>([])
  const [loading, setLoading] = useState(true)
  const [searchFilters, setSearchFilters] = useState({
    city: '',
    type: '',
    minPrice: '',
    maxPrice: ''
  })

  useEffect(() => {
    fetchProperties()
  }, [])

  const fetchProperties = async () => {
    try {
      setLoading(true)
      const data = await propertyApi.getProperties({
        city: searchFilters.city || undefined,
        type: searchFilters.type || undefined,
      })
      setProperties(data)
    } catch (error) {
      console.error('Failed to fetch properties:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = () => {
    fetchProperties()
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <div className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">房源列表</h1>
              <p className="text-gray-600 mt-1">找到您的理想住所</p>
            </div>
            <Button asChild>
              <Link href="/dashboard">返回仪表板</Link>
            </Button>
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Search Filters */}
        <Card className="mb-8">
          <CardHeader>
            <CardTitle>搜索房源</CardTitle>
            <CardDescription>使用筛选条件找到合适的房源</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
              <div>
                <label className="block text-sm font-medium mb-2">城市</label>
                <Input
                  placeholder="输入城市名称"
                  value={searchFilters.city}
                  onChange={(e) => setSearchFilters({...searchFilters, city: e.target.value})}
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">类型</label>
                <select
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  value={searchFilters.type}
                  onChange={(e) => setSearchFilters({...searchFilters, type: e.target.value})}
                >
                  <option value="">所有类型</option>
                  <option value="apartment">公寓</option>
                  <option value="house">别墅</option>
                  <option value="room">单间</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">最低价格</label>
                <Input
                  type="number"
                  placeholder="最低价格"
                  value={searchFilters.minPrice}
                  onChange={(e) => setSearchFilters({...searchFilters, minPrice: e.target.value})}
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">最高价格</label>
                <Input
                  type="number"
                  placeholder="最高价格"
                  value={searchFilters.maxPrice}
                  onChange={(e) => setSearchFilters({...searchFilters, maxPrice: e.target.value})}
                />
              </div>
            </div>
            <div className="mt-4">
              <Button onClick={handleSearch}>搜索房源</Button>
            </div>
          </CardContent>
        </Card>

        {/* Properties Grid */}
        {loading ? (
          <div className="text-center py-12">
            <div className="text-lg text-gray-600">加载中...</div>
          </div>
        ) : properties.length === 0 ? (
          <div className="text-center py-12">
            <div className="text-lg text-gray-600">没有找到符合条件的房源</div>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {properties.map((property) => (
              <Card key={property.id} className="overflow-hidden hover:shadow-lg transition-shadow">
                <div className="aspect-w-16 aspect-h-9 bg-gray-200">
                  {property.images && property.images.length > 0 ? (
                    <img
                      src={property.images[0]}
                      alt={property.title}
                      className="w-full h-48 object-cover"
                    />
                  ) : (
                    <div className="w-full h-48 flex items-center justify-center bg-gray-100">
                      <span className="text-gray-400">暂无图片</span>
                    </div>
                  )}
                </div>
                <CardContent className="p-4">
                  <div className="flex justify-between items-start mb-2">
                    <h3 className="text-lg font-semibold line-clamp-1">{property.title}</h3>
                    <span className={`px-2 py-1 rounded text-xs ${
                      property.available 
                        ? 'bg-green-100 text-green-800' 
                        : 'bg-red-100 text-red-800'
                    }`}>
                      {property.available ? '可租' : '已租'}
                    </span>
                  </div>
                  <p className="text-gray-600 text-sm mb-2">{property.location.address}</p>
                  <p className="text-gray-500 text-sm mb-3 line-clamp-2">{property.description}</p>
                  <div className="flex justify-between items-center">
                    <div>
                      <span className="text-2xl font-bold text-blue-600">¥{property.price}</span>
                      <span className="text-gray-500 text-sm">/月</span>
                    </div>
                    <Button asChild size="sm">
                      <Link href={`/properties/${property.id}`}>查看详情</Link>
                    </Button>
                  </div>
                  {property.features && property.features.length > 0 && (
                    <div className="mt-3 flex flex-wrap gap-1">
                      {property.features.slice(0, 3).map((feature, index) => (
                        <span
                          key={index}
                          className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded"
                        >
                          {feature}
                        </span>
                      ))}
                      {property.features.length > 3 && (
                        <span className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded">
                          +{property.features.length - 3}
                        </span>
                      )}
                    </div>
                  )}
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}
