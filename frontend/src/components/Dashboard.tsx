import React, { useEffect, useState } from 'react';
import { Activity, Calendar, Clock, Target, TrendingUp, Waves, Heart, Flame } from 'lucide-react';
import StatCard from './StatCard';

const Dashboard: React.FC = () => {
    const userId = "59f4e428-9d42-4af8-a18d-1e1dabef47e0";
    const [weeklyStats, setWeeklyStats] = useState({
        distance: 0,
        time: 0,
        avgPace: '0:00',
        avgHeartRate: 0,
        targetDistance: 4000,
        targetTime: 5.0
    });
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchWeeklyStats = async () => {
            setLoading(true);
            try {
                const response = await fetch(`http://localhost:8080/users/${userId}/activities`);
                const data = await response.json();
                const activities = Array.isArray(data.activities) ? data.activities : [];
                // Get start of week (Monday)
                const now = new Date();
                const startOfWeek = new Date(now);
                startOfWeek.setDate(now.getDate() - ((now.getDay() + 6) % 7));
                startOfWeek.setHours(0,0,0,0);
                // Filter activities for this week
                const weekly = activities.filter((a: any) => {
                    const d = new Date(a.date || a.start);
                    return d >= startOfWeek && d <= now;
                });
                const totalDistance = weekly.reduce((sum: number, a: any) => sum + (a.distance || 0), 0);
                const totalTimeMin = weekly.reduce((sum: number, a: any) => {
                    // Parse duration string like '75m', '1h 15m', etc.
                    let min = 0;
                    if (typeof a.duration === 'string') {
                        const hMatch = a.duration.match(/(\d+)h/);
                        const mMatch = a.duration.match(/(\d+)m/);
                        if (hMatch) min += parseInt(hMatch[1], 10) * 60;
                        if (mMatch) min += parseInt(mMatch[1], 10);
                        if (!hMatch && !mMatch && /^\d+$/.test(a.duration)) min = parseInt(a.duration, 10);
                    } else if (typeof a.duration === 'number') {
                        min = a.duration;
                    }
                    return sum + min;
                }, 0);
                // Average pace (min/100m)
                let avgPace = '0:00';
                if (totalDistance > 0 && totalTimeMin > 0) {
                    const pacePerHundred = (totalTimeMin * 60) / (totalDistance / 100);
                    const min = Math.floor(pacePerHundred / 60);
                    const sec = Math.floor(pacePerHundred % 60);
                    avgPace = `${min}:${sec.toString().padStart(2, '0')}`;
                }
                // Average heart rate
                let heartRates: number[] = [];
                weekly.forEach((a: any) => {
                    if (a.heart_rate_avg) heartRates.push(a.heart_rate_avg);
                });
                const avgHeartRate = heartRates.length > 0 ? Math.round(heartRates.reduce((a, b) => a + b, 0) / heartRates.length) : 0;
                setWeeklyStats({
                    distance: totalDistance,
                    time: +(totalTimeMin / 60).toFixed(1),
                    avgPace,
                    avgHeartRate,
                    targetDistance: 4000,
                    targetTime: 5.0
                });
            } catch (err) {
                setWeeklyStats({
                    distance: 0,
                    time: 0,
                    avgPace: '0:00',
                    avgHeartRate: 0,
                    targetDistance: 4000,
                    targetTime: 5.0
                });
            } finally {
                setLoading(false);
            }
        };
        fetchWeeklyStats();
    }, [userId]);

    const stats = [
        {
            title: 'Distância na Semana',
            value: loading ? '...' : weeklyStats.distance.toLocaleString(),
            target: weeklyStats.targetDistance.toLocaleString(),
            progress: weeklyStats.targetDistance ? Math.round((weeklyStats.distance / weeklyStats.targetDistance) * 100) : 0,
            icon: Waves,
            color: 'blue',
            trend: '',
            unit: 'm'
        },
        {
            title: 'Tempo de Natação',
            value: loading ? '...' : weeklyStats.time,
            target: weeklyStats.targetTime,
            progress: weeklyStats.targetTime ? Math.round((weeklyStats.time / weeklyStats.targetTime) * 100) : 0,
            icon: Clock,
            color: 'emerald',
            trend: '',
            unit: 'hrs'
        },
        {
            title: 'Ritmo Médio',
            value: loading ? '...' : weeklyStats.avgPace,
            unit: '/100m',
            icon: TrendingUp,
            color: 'orange',
            trend: ''
        },
        {
            title: 'Frequência Cardíaca Média',
            value: loading ? '...' : weeklyStats.avgHeartRate,
            unit: 'bpm',
            icon: Heart,
            color: 'red',
            trend: ''
        }
    ];

    return (
        <div className="space-y-8">
            {/* Welcome Section */}
            <div className="bg-gradient-to-r from-blue-500 to-cyan-500 rounded-2xl p-6 text-white">
                <div className="flex items-center justify-between">
                    <div>
                        <h2 className="text-2xl font-bold mb-2">Preparado para as atividades de hoje?</h2>
                        {weeklyStats.targetDistance - weeklyStats.distance > 0 ? (
                            <p className="text-blue-100">Você está a {weeklyStats.targetDistance - weeklyStats.distance}m da sua meta semanal</p>
                        ) : (
                            <p className="text-blue-100">Você bateu sua meta!</p>
                        )}
                    </div>
                    <div className="bg-white/20 p-3 rounded-xl backdrop-blur-sm">
                        <Waves className="h-8 w-8" />
                    </div>
                </div>
            </div>

            {/* Stats Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                {stats.map((stat, index) => (
                    <StatCard key={index} {...stat} />
                ))}
            </div>

            
        </div>
    );
};

export default Dashboard;