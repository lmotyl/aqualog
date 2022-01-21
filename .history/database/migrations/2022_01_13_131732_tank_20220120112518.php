<?php

use App\Models\Tank as ModelsTank;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class Tank extends Migration
{

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create(ModelsTank::TABLE_NAME, function (Blueprint $table) {
            $table->id();
            $table->foreignId(ModelsTank::USER_ID);
            $table->foreignId(ModelsTank::TANK_TYPE_ID);
            $table->string(ModelsTank::NAME)->nullable();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists(ModelsTank::TABLE_NAME);
    }
}
