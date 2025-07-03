import React from 'react';
import { DivideIcon as LucideIcon, TrendingUp, TrendingDown } from 'lucide-react';

interface StatCardProps {
    title: string;
    value: string;
    target?: string;
    progress?: number;
    icon: LucideIcon;
    color: 'emerald' | 'orange' | 'blue' | 'red';
    trend?: string;
    unit?: string;
}

const StatCard: React.FC<StatCardProps> = ({ 
    title, 
    value, 
    target, 
    progress, 
    icon: Icon, 
    color, 
    trend,
    unit 
}) => {
    const colorClasses = {
        emerald: 'bg-emerald-500 text-emerald-600 bg-emerald-50',
        orange: 'bg-orange-500 text-orange-600 bg-orange-50',
        blue: 'bg-blue-500 text-blue-600 bg-blue-50',
        red: 'bg-red-500 text-red-600 bg-red-50'
    };

    const [bgColor, textColor, lightBg] = colorClasses[color].split(' ');
    const isPositiveTrend = trend?.startsWith('+');

    return (
        <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200 hover:shadow-md transition-shadow">
            <div className="flex items-center justify-between mb-4">
                <div className={`p-3 rounded-xl ${lightBg}`}>
                    <Icon className={`h-6 w-6 ${textColor}`} />
                </div>
                {trend && (
                    <div className={`flex items-center space-x-1 text-sm font-medium ${
                        isPositiveTrend ? 'text-emerald-600' : 'text-red-500'
                    }`}>
                        {isPositiveTrend ? <TrendingUp className="h-4 w-4" /> : <TrendingDown className="h-4 w-4" />}
                        <span>{trend}</span>
                    </div>
                )}
            </div>
            
            <div className="space-y-2">
                <h3 className="text-slate-600 text-sm font-medium">{title}</h3>
                <div className="flex items-baseline space-x-2">
                    <span className="text-3xl font-bold text-slate-900">{value}</span>
                    {unit && <span className="text-slate-500 text-sm">{unit}</span>}
                </div>
                
                {target && progress && (
                    <div className="space-y-2">
                        <div className="flex justify-between text-sm">
                            <span className="text-slate-500">Meta: {target}</span>
                            <span className={`font-medium ${textColor}`}>{progress}%</span>
                        </div>
                        <div className="w-full bg-slate-200 rounded-full h-2">
                            <div 
                                className={`${bgColor} h-2 rounded-full transition-all duration-500 ease-out`}
                                style={{ width: `${Math.min(progress, 100)}%` }}
                            />
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};

export default StatCard;