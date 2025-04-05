import React, { useState } from 'react';
import { ShoppingCart, Search, Menu, X, Heart } from 'lucide-react';
import ProductCard from './components/ProductCard';
import CartSidebar from './components/CartSidebar';
import { Product } from './types';

function App() {
  const [isCartOpen, setIsCartOpen] = useState(false);
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  // Mock products data
  const products: Product[] = [
    {
      id: 1,
      name: "Premium Wireless Headphones",
      price: 299.99,
      description: "High-quality wireless headphones with noise cancellation",
      image: "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=500&q=80",
      category: "Electronics"
    },
    {
      id: 2,
      name: "Smart Watch Pro",
      price: 199.99,
      description: "Advanced smartwatch with health tracking features",
      image: "https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=500&q=80",
      category: "Electronics"
    },
    {
      id: 3,
      name: "Premium Leather Wallet",
      price: 79.99,
      description: "Handcrafted genuine leather wallet",
      image: "https://images.unsplash.com/photo-1627123424574-724758594e93?w=500&q=80",
      category: "Accessories"
    },
    {
      id: 4,
      name: "Minimalist Backpack",
      price: 129.99,
      description: "Stylish and functional everyday backpack",
      image: "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=500&q=80",
      category: "Accessories"
    }
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Navigation */}
      <nav className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <button
                className="sm:hidden p-2"
                onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
              >
                {isMobileMenuOpen ? <X size={24} /> : <Menu size={24} />}
              </button>
              <span className="text-2xl font-bold text-indigo-600">ShopHub</span>
            </div>

            {/* Desktop Navigation */}
            <div className="hidden sm:flex items-center space-x-8">
              <a href="#" className="text-gray-600 hover:text-gray-900">Home</a>
              <a href="#" className="text-gray-600 hover:text-gray-900">Shop</a>
              <a href="#" className="text-gray-600 hover:text-gray-900">Categories</a>
              <a href="#" className="text-gray-600 hover:text-gray-900">About</a>
            </div>

            <div className="flex items-center space-x-4">
              <button className="p-2">
                <Search size={24} className="text-gray-600" />
              </button>
              <button className="p-2">
                <Heart size={24} className="text-gray-600" />
              </button>
              <button 
                className="p-2 relative"
                onClick={() => setIsCartOpen(true)}
              >
                <ShoppingCart size={24} className="text-gray-600" />
                <span className="absolute top-0 right-0 bg-indigo-600 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs">
                  3
                </span>
              </button>
            </div>
          </div>
        </div>

        {/* Mobile Navigation */}
        {isMobileMenuOpen && (
          <div className="sm:hidden bg-white border-t">
            <div className="px-2 pt-2 pb-3 space-y-1">
              <a href="#" className="block px-3 py-2 text-gray-600 hover:text-gray-900">Home</a>
              <a href="#" className="block px-3 py-2 text-gray-600 hover:text-gray-900">Shop</a>
              <a href="#" className="block px-3 py-2 text-gray-600 hover:text-gray-900">Categories</a>
              <a href="#" className="block px-3 py-2 text-gray-600 hover:text-gray-900">About</a>
            </div>
          </div>
        )}
      </nav>

      {/* Hero Section */}
      <div className="relative bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
          <div className="text-center">
            <h1 className="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl">
              <span className="block">Premium Quality</span>
              <span className="block text-indigo-600">Products for Everyone</span>
            </h1>
            <p className="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
              Discover our curated collection of premium products. From electronics to accessories, 
              we've got everything you need.
            </p>
            <div className="mt-5 max-w-md mx-auto sm:flex sm:justify-center md:mt-8">
              <div className="rounded-md shadow">
                <a href="#" className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 md:py-4 md:text-lg md:px-10">
                  Shop Now
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Products Grid */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-8">Featured Products</h2>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {products.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      </div>

      {/* Cart Sidebar */}
      <CartSidebar isOpen={isCartOpen} onClose={() => setIsCartOpen(false)} />

      {/* Footer */}
      <footer className="bg-gray-800">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
            <div>
              <h3 className="text-white text-lg font-bold mb-4">Ecom Platform</h3>
            </div>
            
            <div>
              <h4 className="text-white text-lg font-bold mb-4">Categories</h4>
              <ul className="space-y-2">
                <li><a href="#" className="text-gray-400 hover:text-white">Electronics</a></li>
                <li><a href="#" className="text-gray-400 hover:text-white">Accessories</a></li>
                <li><a href="#" className="text-gray-400 hover:text-white">Clothing</a></li>
              </ul>
            </div>
            
          </div>
        </div>
      </footer>
    </div>
  );
}

export default App;