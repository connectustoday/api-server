import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class INotification {

	@prop({ required: true })
	id: number;

	@prop({ required: true })
	created_at: number;

	@prop({ required: true })
	type: string;

	@prop({ required: true })
	content: string;

	@prop()
	account?: string;
}
