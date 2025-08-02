import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import Link from 'next/link'
import { Button } from '@/components/ui/button'

export default function AboutPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Navigation */}
      <nav className="border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <div className="container flex h-14 items-center">
          <div className="mr-4 flex">
            <Link href="/" className="mr-6 flex items-center space-x-2">
              <span className="font-bold text-xl">RentHelp</span>
            </Link>
          </div>
          <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
            <nav className="flex items-center space-x-6">
              <Link href="/properties" className="text-sm font-medium">
                房源
              </Link>
              <Link href="/about" className="text-sm font-medium text-blue-600">
                关于我们
              </Link>
            </nav>
            <div className="flex items-center space-x-2">
              <Button variant="ghost" asChild>
                <Link href="/login">登录</Link>
              </Button>
              <Button asChild>
                <Link href="/register">注册</Link>
              </Button>
            </div>
          </div>
        </div>
      </nav>

      <div className="container mx-auto px-4 py-12">
        <div className="max-w-4xl mx-auto">
          {/* Hero Section */}
          <div className="text-center mb-12">
            <h1 className="text-4xl font-bold text-gray-900 mb-4">
              关于 RentHelp
            </h1>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              我们致力于打造最专业、最便捷的租房服务平台，让每个人都能找到理想的住所。
            </p>
          </div>

          {/* Mission Section */}
          <Card className="mb-8">
            <CardHeader>
              <CardTitle className="text-2xl">我们的使命</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-gray-700 leading-relaxed">
                RentHelp 致力于通过技术创新来解决租房市场中的痛点问题。我们相信租房应该是一个透明、高效、安全的过程。
                通过我们的平台，房东可以轻松管理房源，租客可以快速找到满意的住所，双方都能享受到优质的服务体验。
              </p>
            </CardContent>
          </Card>

          {/* Features Grid */}
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">🏠 丰富房源</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  覆盖全城优质房源，从精装公寓到温馨住宅，总有一款适合您。
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">🔍 智能搜索</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  基于AI的智能推荐系统，根据您的需求精准匹配最合适的房源。
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">🛡️ 安全保障</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  严格的房源审核机制和用户认证体系，确保每一次交易都安全可靠。
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">💬 贴心服务</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  专业的客服团队7×24小时在线，随时为您解答疑问，提供帮助。
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">📱 便捷体验</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  响应式设计，支持PC、手机、平板等多设备访问，随时随地找房看房。
                </p>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle className="text-lg">💰 透明定价</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-gray-600">
                  无隐藏费用，所有价格信息透明公开，让您租房更放心。
                </p>
              </CardContent>
            </Card>
          </div>

          {/* Team Section */}
          <Card className="mb-8">
            <CardHeader>
              <CardTitle className="text-2xl">我们的团队</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-gray-700 leading-relaxed mb-4">
                RentHelp 团队由来自房地产、互联网、技术等多个领域的专业人士组成。我们拥有丰富的行业经验和深厚的技术实力，
                致力于为用户提供最优质的租房服务。
              </p>
              <div className="grid md:grid-cols-3 gap-4 mt-6">
                <div className="text-center">
                  <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-2xl">👨‍💼</span>
                  </div>
                  <h4 className="font-semibold">产品团队</h4>
                  <p className="text-sm text-gray-600">专注用户体验设计</p>
                </div>
                <div className="text-center">
                  <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-2xl">👨‍💻</span>
                  </div>
                  <h4 className="font-semibold">技术团队</h4>
                  <p className="text-sm text-gray-600">保障平台稳定运行</p>
                </div>
                <div className="text-center">
                  <div className="w-16 h-16 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-2">
                    <span className="text-2xl">👥</span>
                  </div>
                  <h4 className="font-semibold">服务团队</h4>
                  <p className="text-sm text-gray-600">提供专业客户服务</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Contact Section */}
          <Card>
            <CardHeader>
              <CardTitle className="text-2xl">联系我们</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid md:grid-cols-2 gap-6">
                <div>
                  <h4 className="font-semibold mb-2">📧 联系方式</h4>
                  <p className="text-gray-600 mb-1">邮箱：support@renthelp.com</p>
                  <p className="text-gray-600 mb-1">电话：400-888-8888</p>
                  <p className="text-gray-600">微信：RentHelp客服</p>
                </div>
                <div>
                  <h4 className="font-semibold mb-2">🏢 公司地址</h4>
                  <p className="text-gray-600 mb-1">北京市朝阳区建国门外大街1号</p>
                  <p className="text-gray-600 mb-1">国贸大厦A座20层</p>
                  <p className="text-gray-600">邮编：100020</p>
                </div>
              </div>
              <div className="mt-6 pt-6 border-t">
                <div className="flex flex-wrap gap-4">
                  <Button variant="outline" asChild>
                    <Link href="/properties">浏览房源</Link>
                  </Button>
                  <Button variant="outline" asChild>
                    <Link href="/register">立即注册</Link>
                  </Button>
                  <Button variant="outline" asChild>
                    <Link href="/">返回首页</Link>
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
