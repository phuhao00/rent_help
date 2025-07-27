'use client'

import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { cn } from '@/lib/utils'

interface DateRangePickerProps {
  onDateSelect: (startDate: Date, endDate: Date) => void
  minDate?: Date
  maxDate?: Date
  className?: string
}

export function DateRangePicker({ 
  onDateSelect, 
  minDate = new Date(), 
  maxDate,
  className 
}: DateRangePickerProps) {
  const [startDate, setStartDate] = useState('')
  const [endDate, setEndDate] = useState('')
  const [error, setError] = useState('')

  const formatDate = (date: Date): string => {
    return date.toISOString().split('T')[0]
  }

  const validateDates = (start: string, end: string): boolean => {
    if (!start || !end) {
      setError('请选择开始和结束日期')
      return false
    }

    const startDateObj = new Date(start)
    const endDateObj = new Date(end)

    if (startDateObj >= endDateObj) {
      setError('结束日期必须晚于开始日期')
      return false
    }

    if (startDateObj < minDate) {
      setError('开始日期不能早于今天')
      return false
    }

    if (maxDate && endDateObj > maxDate) {
      setError(`结束日期不能晚于 ${formatDate(maxDate)}`)
      return false
    }

    setError('')
    return true
  }

  const handleDateChange = () => {
    if (validateDates(startDate, endDate)) {
      onDateSelect(new Date(startDate), new Date(endDate))
    }
  }

  const calculateDays = (): number => {
    if (!startDate || !endDate) return 0
    const start = new Date(startDate)
    const end = new Date(endDate)
    const diffTime = Math.abs(end.getTime() - start.getTime())
    return Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  }

  return (
    <Card className={cn('w-full', className)}>
      <CardHeader>
        <CardTitle>选择租期</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium mb-2">开始日期</label>
            <Input
              type="date"
              value={startDate}
              min={formatDate(minDate)}
              max={maxDate ? formatDate(maxDate) : undefined}
              onChange={(e) => {
                setStartDate(e.target.value)
                if (endDate) handleDateChange()
              }}
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">结束日期</label>
            <Input
              type="date"
              value={endDate}
              min={startDate || formatDate(minDate)}
              max={maxDate ? formatDate(maxDate) : undefined}
              onChange={(e) => {
                setEndDate(e.target.value)
                if (startDate) handleDateChange()
              }}
            />
          </div>
        </div>

        {error && (
          <div className="text-sm text-red-600">
            {error}
          </div>
        )}

        {startDate && endDate && !error && (
          <div className="p-3 bg-green-50 border border-green-200 rounded-md">
            <p className="text-sm text-green-800">
              租期: {calculateDays()} 天
            </p>
          </div>
        )}
      </CardContent>
    </Card>
  )
}

// Rating Component
interface RatingProps {
  value: number
  max?: number
  readonly?: boolean
  onRatingChange?: (rating: number) => void
  size?: 'sm' | 'md' | 'lg'
  className?: string
}

export function Rating({ 
  value, 
  max = 5, 
  readonly = false, 
  onRatingChange,
  size = 'md',
  className 
}: RatingProps) {
  const [hoverValue, setHoverValue] = useState(0)

  const sizeClasses = {
    sm: 'w-4 h-4',
    md: 'w-6 h-6',
    lg: 'w-8 h-8'
  }

  const handleClick = (rating: number) => {
    if (!readonly && onRatingChange) {
      onRatingChange(rating)
    }
  }

  const handleMouseEnter = (rating: number) => {
    if (!readonly) {
      setHoverValue(rating)
    }
  }

  const handleMouseLeave = () => {
    if (!readonly) {
      setHoverValue(0)
    }
  }

  return (
    <div className={cn('flex items-center space-x-1', className)}>
      {Array.from({ length: max }, (_, index) => {
        const starValue = index + 1
        const isFilled = starValue <= (hoverValue || value)
        
        return (
          <button
            key={index}
            type="button"
            disabled={readonly}
            className={cn(
              'transition-colors',
              readonly ? 'cursor-default' : 'cursor-pointer hover:scale-110',
              sizeClasses[size]
            )}
            onClick={() => handleClick(starValue)}
            onMouseEnter={() => handleMouseEnter(starValue)}
            onMouseLeave={handleMouseLeave}
          >
            <svg
              className={cn(
                'w-full h-full transition-colors',
                isFilled ? 'text-yellow-400 fill-current' : 'text-gray-300'
              )}
              viewBox="0 0 20 20"
            >
              <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
            </svg>
          </button>
        )
      })}
      <span className="ml-2 text-sm text-gray-600">
        {value.toFixed(1)} / {max}
      </span>
    </div>
  )
}

// Search Filter Component
interface SearchFiltersProps {
  onFiltersChange: (filters: any) => void
  className?: string
}

export function SearchFilters({ onFiltersChange, className }: SearchFiltersProps) {
  const [filters, setFilters] = useState({
    location: '',
    type: '',
    minPrice: '',
    maxPrice: '',
    bedrooms: '',
    bathrooms: '',
    amenities: [] as string[]
  })

  const amenitiesList = [
    'WiFi',
    '空调',
    '洗衣机',
    '冰箱',
    '微波炉',
    '电视',
    '停车位',
    '电梯',
    '健身房',
    '游泳池'
  ]

  const handleFilterChange = (key: string, value: any) => {
    const newFilters = { ...filters, [key]: value }
    setFilters(newFilters)
    onFiltersChange(newFilters)
  }

  const handleAmenityToggle = (amenity: string) => {
    const newAmenities = filters.amenities.includes(amenity)
      ? filters.amenities.filter(a => a !== amenity)
      : [...filters.amenities, amenity]
    
    handleFilterChange('amenities', newAmenities)
  }

  const clearFilters = () => {
    const emptyFilters = {
      location: '',
      type: '',
      minPrice: '',
      maxPrice: '',
      bedrooms: '',
      bathrooms: '',
      amenities: []
    }
    setFilters(emptyFilters)
    onFiltersChange(emptyFilters)
  }

  return (
    <Card className={cn('w-full', className)}>
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle>筛选条件</CardTitle>
        <Button variant="outline" size="sm" onClick={clearFilters}>
          清空筛选
        </Button>
      </CardHeader>
      <CardContent className="space-y-4">
        <div>
          <label className="block text-sm font-medium mb-2">位置</label>
          <Input
            placeholder="输入城市或地区"
            value={filters.location}
            onChange={(e) => handleFilterChange('location', e.target.value)}
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">房屋类型</label>
          <select
            className="w-full p-2 border border-gray-300 rounded-md"
            value={filters.type}
            onChange={(e) => handleFilterChange('type', e.target.value)}
          >
            <option value="">全部类型</option>
            <option value="apartment">公寓</option>
            <option value="house">独栋房屋</option>
            <option value="condo">共管公寓</option>
            <option value="townhouse">联排别墅</option>
          </select>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium mb-2">最低价格</label>
            <Input
              type="number"
              placeholder="0"
              value={filters.minPrice}
              onChange={(e) => handleFilterChange('minPrice', e.target.value)}
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">最高价格</label>
            <Input
              type="number"
              placeholder="无限制"
              value={filters.maxPrice}
              onChange={(e) => handleFilterChange('maxPrice', e.target.value)}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium mb-2">卧室数量</label>
            <select
              className="w-full p-2 border border-gray-300 rounded-md"
              value={filters.bedrooms}
              onChange={(e) => handleFilterChange('bedrooms', e.target.value)}
            >
              <option value="">任意</option>
              <option value="1">1室</option>
              <option value="2">2室</option>
              <option value="3">3室</option>
              <option value="4">4室+</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">卫生间数量</label>
            <select
              className="w-full p-2 border border-gray-300 rounded-md"
              value={filters.bathrooms}
              onChange={(e) => handleFilterChange('bathrooms', e.target.value)}
            >
              <option value="">任意</option>
              <option value="1">1间</option>
              <option value="2">2间</option>
              <option value="3">3间+</option>
            </select>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">设施</label>
          <div className="grid grid-cols-2 gap-2">
            {amenitiesList.map((amenity) => (
              <label key={amenity} className="flex items-center space-x-2">
                <input
                  type="checkbox"
                  checked={filters.amenities.includes(amenity)}
                  onChange={() => handleAmenityToggle(amenity)}
                  className="rounded"
                />
                <span className="text-sm">{amenity}</span>
              </label>
            ))}
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
