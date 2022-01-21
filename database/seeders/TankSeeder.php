<?php

namespace Database\Seeders;

use App\Models\Tank;
use Database\Factories\TankFactory;
use Illuminate\Database\Seeder;

class TankSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {


        Tank::create([ 
            Tank::NAME => 'first small',
            TANK::USER_ID => 1,
            TANK::TANK_TYPE_ID => 1,
        ]);
        Tank::create([ 
            Tank::NAME => 'first small',
            TANK::USER_ID => 2,
            TANK::TANK_TYPE_ID => 1,
        ]);
        Tank::create([ 
            Tank::NAME => 'second medium',
            TANK::USER_ID => 2,
            TANK::TANK_TYPE_ID => 2,
        ]);
        Tank::create([ 
            Tank::NAME => 'third big',
            TANK::USER_ID => 2,
            TANK::TANK_TYPE_ID => 3,
        ]);
    }
}
