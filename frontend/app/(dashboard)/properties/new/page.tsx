'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { propertyApi } from '@/lib/api'
import { useAuthStore } from '@/store/auth'

interface PropertyForm {
  title: string
  description: string
  type: 'apartment' | 'house' | 'room'
  price: string
  currency: string
  location: {
    address: string
    city: string
    state: string
    country: string
    latitude: number
    longitude: number
  }
  features: string[]
  images: string[]
}

export default function NewPropertyPage() {
  const router = useRouter()
  const { user } = useAuthStore()
  const [loading, setLoading] = useState(false)
  const [featureInput, setFeatureInput] = useState('')
  const [imageInput, setImageInput] = useState('')
  
  const [formData, setFormData] = useState<PropertyForm>({
    title: '',
    description: '',
    type: 'apartment',
    price: '',
    currency: 'CNY',
    location: {
      address: '',
      city: '',
      state: '',
      country: '中国',
      latitude: 0,
      longitude: 0
    },
    features: [],
    images: []
  })

  const handleInputChange = (field: string, value: any, nested?: string) => {
    if (nested) {
      setFormData(prev => ({
        ...prev,
        [nested]: {
          ...prev[nested as keyof PropertyForm],
          [field]: value
        }
      }))
    } else {
      setFormData(prev => ({
        ...prev,
        [field]: value
      }))
    }
  }

  const addFeature = () => {
    if (featureInput.trim()) {
      setFormData(prev => ({
        ...prev,
        features: [...prev.features, featureInput.trim()]
      }))
      setFeatureInput('')
    }
  }

  const removeFeature = (index: number) => {
    setFormData(prev => ({
      ...prev,
      features: prev.features.filter((_, i) => i !== index)
    }))
  }

  const addImage = () => {
    if (imageInput.trim()) {
      setFormData(prev => ({
        ...prev,
        images: [...prev.images, imageInput.trim()]
      }))
      setImageInput('')
    }
  }

  const removeImage = (index: number) => {
    setFormData(prev => ({
      ...prev,
      images: prev.images.filter((_, i) => i !== index)
    }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    if (user?.role !== 'landlord') {
      alert('只有房东可以发布房源')
      return
    }

    setLoading(true)
    try {
      const propertyData = {
        ...formData,
        price: parseFloat(formData.price),
        available: true
      }

      await propertyApi.createProperty(propertyData)
      alert('房源发布成功！')
      router.push('/dashboard')
    } catch (error) {
      console.error('Failed to create property:', error)
      alert('发布失败，请重试')
    } finally {
      setLoading(false)
    }
  }

  if (user?.role !== 'landlord') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <Card className="w-full max-w-md">
          <CardContent className="pt-6">
            <div className="text-center">
              <h2 className="text-xl font-semibold mb-4">权限不足</h2>
              <p className="text-gray-600 mb-4">只有房东可以发布房源</p>
              <Button onClick={() => router.push('/dashboard')}>
                返回仪表板
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-6">
          <Button variant="outline" onClick={() => router.back()}>
            ← 返回
          </Button>
        </div>

        <Card>
          <CardHeader>
            <CardTitle>发布新房源</CardTitle>
            <CardDescription>
              填写详细信息来发布您的房源
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleSubmit} className="space-y-6">
              {/* 基本信息 */}
              <div className="space-y-4">
                <h3 className="text-lg font-semibold">基本信息</h3>
                
                <div>
                  <label className="block text-sm font-medium mb-2">房源标题</label>
                  <Input
                    value={formData.title}
                    onChange={(e) => handleInputChange('title', e.target.value)}
                    placeholder="输入吸引人的房源标题"
                    required
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium mb-2">房源描述</label>
                  <textarea
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    rows={4}
                    value={formData.description}
                    onChange={(e) => handleInputChange('description', e.target.value)}
                    placeholder="详细描述您的房源特色、周边环境等"
                    required
                  />
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label className="block text-sm font-medium mb-2">房源类型</label>
                    <select
                      className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                      value={formData.type}
                      onChange={(e) => handleInputChange('type', e.target.value)}
                    >
                      <option value="apartment">公寓</option>
                      <option value="house">别墅</option>
                      <option value="room">单间</option>
                    </select>
                  </div>

                  <div>
                    <label className="block text-sm font-medium mb-2">租金（元/月）</label>
                    <Input
                      type="number"
                      value={formData.price}
                      onChange={(e) => handleInputChange('price', e.target.value)}
                      placeholder="输入月租金"
                      required
                      min="0"
                    />
                  </div>
                </div>
              </div>

              {/* 位置信息 */}
              <div className="space-y-4">
                <h3 className="text-lg font-semibold">位置信息</h3>
                
                <div>
                  <label className="block text-sm font-medium mb-2">详细地址</label>
                  <Input
                    value={formData.location.address}
                    onChange={(e) => handleInputChange('address', e.target.value, 'location')}
                    placeholder="街道地址、门牌号等"
                    required
                  />
                </div>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                  <div>
                    <label className="block text-sm font-medium mb-2">城市</label>
                    <Input
                      value={formData.location.city}
                      onChange={(e) => handleInputChange('city', e.target.value, 'location')}
                      placeholder="城市"
                      required
                    />
                  </div>

                  <div>
                    <label className="block text-sm font-medium mb-2">省份</label>
                    <Input
                      value={formData.location.state}
                      onChange={(e) => handleInputChange('state', e.target.value, 'location')}
                      placeholder="省份"
                      required
                    />
                  </div>

                  <div>
                    <label className="block text-sm font-medium mb-2">国家</label>
                    <Input
                      value={formData.location.country}
                      onChange={(e) => handleInputChange('country', e.target.value, 'location')}
                      placeholder="国家"
                      required
                    />
                  </div>
                </div>
              </div>

              {/* 房源特色 */}
              <div className="space-y-4">
                <h3 className="text-lg font-semibold">房源特色</h3>
                
                <div className="flex gap-2">
                  <Input
                    value={featureInput}
                    onChange={(e) => setFeatureInput(e.target.value)}
                    placeholder="输入房源特色（如：WiFi、空调、停车位等）"
                    onKeyPress={(e) => e.key === 'Enter' && (e.preventDefault(), addFeature())}
                  />
                  <Button type="button" onClick={addFeature}>
                    添加
                  </Button>
                </div>

                {formData.features.length > 0 && (
                  <div className="flex flex-wrap gap-2">
                    {formData.features.map((feature, index) => (
                      <div
                        key={index}
                        className="flex items-center gap-1 px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-sm"
                      >
                        {feature}
                        <button
                          type="button"
                          onClick={() => removeFeature(index)}
                          className="ml-1 text-blue-600 hover:text-blue-800"
                        >
                          ×
                        </button>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              {/* 房源图片 */}
              <div className="space-y-4">
                <h3 className="text-lg font-semibold">房源图片</h3>
                
                <div className="flex gap-2">
                  <Input
                    value={imageInput}
                    onChange={(e) => setImageInput(e.target.value)}
                    placeholder="输入图片URL"
                    onKeyPress={(e) => e.key === 'Enter' && (e.preventDefault(), addImage())}
                  />
                  <Button type="button" onClick={addImage}>
                    添加
                  </Button>
                </div>

                {formData.images.length > 0 && (
                  <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
                    {formData.images.map((image, index) => (
                      <div key={index} className="relative">
                        <img
                          src={image}
                          alt={`房源图片 ${index + 1}`}
                          className="w-full h-32 object-cover rounded-lg"
                        />
                        <button
                          type="button"
                          onClick={() => removeImage(index)}
                          className="absolute top-2 right-2 bg-red-500 text-white rounded-full w-6 h-6 flex items-center justify-center text-sm hover:bg-red-600"
                        >
                          ×
                        </button>
                      </div>
                    ))}
                  </div>
                )}
              </div>

              <div className="flex gap-4 pt-6">
                <Button type="submit" disabled={loading} className="flex-1">
                  {loading ? '发布中...' : '发布房源'}
                </Button>
                <Button type="button" variant="outline" onClick={() => router.back()}>
                  取消
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
